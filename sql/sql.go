package sql

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	rc := filepath.Join(os.Getenv("HOMEPATH"), "sqltool.rc")
	// TODO: Add mac.
	if _, err := os.Stat(rc); os.IsNotExist(err) {
		file, err := os.Create(rc)
		if err != nil {
			panic(err)
		}
		file.Write([]byte("urlid skirmishdb\n" +
			"url jdbc:hsqldb:file:///F:\\Downloads\\SkirmishDB\\database/skirmishdb;" +
			"default_schema=true;shutdown=true;hsqldb.default_table_type=cached;get_column_name=false\n" +
			"username sa\n" +
			"password\n"))
	}
}

func Query(query string) ([]byte, error) {
	var out bytes.Buffer
	var errs bytes.Buffer
	cmd := exec.Command("java", "-jar", ".\\sqltool.jar", "--noinput",
		"--sql", query, "skirmishdb")
	cmd.Stdout = &out
	cmd.Stderr = &errs
	err := cmd.Run()
	if err != nil {
		return errs.Bytes(), err
	}
	return out.Bytes(), nil
}
