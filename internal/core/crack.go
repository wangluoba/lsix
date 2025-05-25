package core

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	addOpens1 = "--add-opens=java.base/jdk.internal.org.objectweb.asm=ALL-UNNAMED"
	addOpens2 = "--add-opens=java.base/jdk.internal.org.objectweb.asm.tree=ALL-UNNAMED"
	agentLine = "-javaagent:" + jaNetfilter + "=jetbrains"
)

func setUserConfigPath() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "JetBrains")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "JetBrains")
	default: // Linux
		return filepath.Join(os.Getenv("HOME"), ".config", "JetBrains")
	}
}

type CrackRequest struct {
	App     string  `json:"app"`
	Status  string  `json:"status"`
	License License `json:"license"`
}

func CrackHandler(c *gin.Context) {
	var req CrackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ Parsing request failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app := req.App
	status := req.Status
	log.Printf("🚀 Start processing the request: app=%s, status=%s", app, status)
	action := ""
	if app == "IntelliJ IDEA" {
		app = "IntelliJIdea"
	}
	currentStatus := GetCrackStatus(app)

	path := setUserConfigPath()

	dirs, err := os.ReadDir(path)
	if err != nil {
		log.Printf("❌ Failed to read directory: %v", err)
		c.JSON(500, gin.H{
			"error":  "Failed to read directory",
			"detail": err.Error(),
		})
		return
	}
	backup := false
	setkey := false
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		if !strings.HasPrefix(dir.Name(), app) {
			continue
		}
		dirPath := filepath.Join(path, dir.Name())
		// vmoptionsFiles
		vmoptionsFiles, _ := filepath.Glob(filepath.Join(dirPath, "*.vmoptions"))
		// keyfile
		keyfile := filepath.Join(dirPath, strings.ToLower(app)+".key")
		if strings.HasSuffix(keyfile, "intellijidea.key") {
			keyfile = filepath.Join(dirPath, "idea.key")
		}

		if status == "Cracked" {

			for _, file := range vmoptionsFiles {
				backupPath := file + ".jetbra-free.bak"
				if currentStatus == "UnCracked" {
					data, err := os.ReadFile(file)
					if err == nil {
						log.Printf("📦 Backup vmoptions: %s", backupPath)
						_ = os.WriteFile(backupPath, data, 0644)
					}
					err = editVmoptionsFile(file)
					if err != nil {
						fmt.Printf("Failed to patch %s: %v\n", file, err)
					}
					backup = true
				}

			}

			licenseStr, _ := GenerateLicense(&req.License)

			if currentStatus == "UnCracked" {
				keyBackupPath := keyfile + ".jetbra-free.bak"
				data, err := os.ReadFile(keyfile)
				if err == nil {
					_ = os.WriteFile(keyBackupPath, data, 0644)
					log.Printf("🔑 Backup key: %s", keyBackupPath)
				}
			}

			if err := setKeyFile(keyfile, licenseStr); err != nil {
				log.Printf("❌ Failed to set key file: %v", err)
			}
			setkey = true
			if setkey && !backup {
				action = "CrackedWithoutBackup"
			} else {
				action = "Cracked"
			}

		} else if status == "UnCracked" {
			for _, file := range vmoptionsFiles {
				if err := revertVmoptionsFile(file); err != nil {
					log.Printf("❌ Restore failed: %v", err)
				}
			}

			if err := revertKeyFile(keyfile); err != nil {
				log.Printf("❌ Failed to restore key file: %v", err)
			}
			action = "UnCracked"
		}
	}

	if action != "" {
		c.JSON(200, gin.H{"msg": action})
		return
	}
	c.JSON(200, gin.H{"msg": "ERROR"})
}

func setKeyFile(keyPath string, licenseStr string) error {
	content := "\uFFFF<certificate-key>\n" + licenseStr

	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	utf16Content, _, err := transform.String(encoder, content)
	if err != nil {
		return err
	}
	log.Printf("🔑 Write key: %s", keyPath)
	return os.WriteFile(keyPath, []byte(utf16Content), 0644)
}

func revertKeyFile(keyPath string) error {

	backupPath := keyPath + ".jetbra-free.bak"
	if _, err := os.Stat(backupPath); err == nil {
		err := os.Rename(backupPath, keyPath)
		if err != nil {
			return err
		}
		log.Printf("🔑 Revert key: %s", keyPath)
		return nil
	}
	if _, err := os.Stat(keyPath); err == nil {
		log.Printf("🔑 Delete key: %s", keyPath)
		return os.Remove(keyPath)
	}
	return nil
}

func editVmoptionsFile(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		lines        []string
		hasAddOpens1 bool
		hasAddOpens2 bool
		hasJavaAgent bool
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, addOpens1) {
			hasAddOpens1 = true
		}
		if strings.Contains(line, addOpens2) {
			hasAddOpens2 = true
		}
		if strings.HasPrefix(line, "-javaagent:") && strings.Contains(line, "=jetbrains") {
			hasJavaAgent = true
			lines = append(lines, agentLine)
			continue
		}
		lines = append(lines, line)

	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if !hasAddOpens1 {
		lines = append(lines, addOpens1)
	}
	if !hasAddOpens2 {
		lines = append(lines, addOpens2)
	}
	if !hasJavaAgent {
		lines = append(lines, agentLine)
	}

	log.Printf("📄 Edit vmoptions: %s", path)
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

func revertVmoptionsFile(path string) error {
	backupPath := path + ".jetbra-free.bak"

	_, err := os.Stat(backupPath)
	if err == nil {
		data, err := os.ReadFile(backupPath)
		if err != nil {
			println("Error reading backup file:", err)
			return err
		}
		if err := os.WriteFile(path, data, 0644); err != nil {
			println("Error writing to vmoptions file:", err)
			return err
		}
		log.Println("📄 Revert vmoptions:", path)
		return os.Remove(backupPath)
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trim := strings.TrimSpace(line)
		if trim == addOpens1 {
			continue
		}
		if trim == addOpens2 {
			continue
		}
		if strings.HasPrefix(trim, "-javaagent:") && strings.Contains(trim, "=jetbrains") {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	log.Printf("📄 Remove the hack in the vmoptions file: %s", path)
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

func GetCrackStatus(appPrefix string) string {
	app := appPrefix
	if app == "IntelliJIdea" {
		app = "IntelliJ IDEA"
	}

	path := setUserConfigPath()
	dirs, err := os.ReadDir(path)
	if err != nil {
		log.Printf("❌ Failed to read directory: %v\n", err)
		return "error"
	}

	cracked := true
	found := false

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		if !strings.HasPrefix(dir.Name(), appPrefix) {
			continue
		}
		found = true

		dirPath := filepath.Join(path, dir.Name())
		pattern := "*.vmoptions"
		files, _ := filepath.Glob(filepath.Join(dirPath, pattern))

		for _, file := range files {
			data, err := os.ReadFile(file)
			if err != nil {
				cracked = false
				continue
			}
			content := string(data)
			lines := strings.Split(content, "\n")
			hasAddOpens1 := false
			hasAddOpens2 := false
			hasAgentLine := false
			for _, line := range lines {
				trim := strings.TrimSpace(line)
				if trim == addOpens1 {
					hasAddOpens1 = true
				}
				if trim == addOpens2 {
					hasAddOpens2 = true
				}
				if trim == agentLine {
					hasAgentLine = true
				}
			}
			if !(hasAddOpens1 && hasAddOpens2 && hasAgentLine) {
				cracked = false
				log.Printf("🔒 %s Current Status: UnCracked 📄 %s\n", app, file)
			} else {
				log.Printf("🎉 %s Current Status: Cracked 📄 %s\n", app, file)

			}
		}
	}

	if !found {
		fmt.Printf("ℹ️  %s: Uninstall \n", app)
		return "Uninstall"
	}
	if cracked {
		return "Cracked"
	}
	return "UnCracked"
}
