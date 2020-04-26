package gradexpath

import (
	"os"
	"path/filepath"
)

var (
	isTesting bool
	testroot  = "./tmp-delete-me"
	ExamStage = []string{
		"00-config",
		"10-temp-receipt",
		"11-temp-pdf",
		"12-temp-annotate",
		"13-temp-reject",
		"14-temp-paper",
		"15-temp-mark",
		"16-temp-moderate",
		"17-temp-check",
		"18-temp-config",
		"19-temp-ignore",
		"20-partial-paper-set",
		"21-patch-set",
		"22-complete-paper-set",
		"40-ready-to-mark",
		"42-already-sent-to-marker",
		"50-from-marker",
		"58-marker-merged",
		"60-to-moderator",
		"70-from-moderator",
		"78-moderator-merged",
		"80-to-checker",
		"88-checker-merged",
		"90-from-checker",
		"99-reports",
	}
)

func Ingest() string {
	return filepath.Join(Root(), "ingest")
}

func Export() string {
	return filepath.Join(Root(), "export")
}

func Etc() string {
	return filepath.Join(Root(), "etc")
}

func Var() string {
	return filepath.Join(Root(), "var")
}

func Usr() string {
	return filepath.Join(Root(), "usr")
}

func Exam() string {
	return filepath.Join(Usr(), "exam")
}

func IngestConf() string {
	return filepath.Join(Etc(), "ingest")
}

func OverlayConf() string {
	return filepath.Join(Etc(), "overlay")
}

func ExtractConf() string {
	return filepath.Join(Etc(), "extract")
}

func SetupConf() string {
	return filepath.Join(Etc(), "setup")
}

func SetTesting() { //need this when testing other tools
	isTesting = true
}

func Root() string {
	if isTesting {
		return testroot
	}
	return root
}

func GetExamPath(name string) string {
	return filepath.Join(Exam(), name)
}

func GetExamStagePath(name, stage string) string {
	return filepath.Join(Exam(), name, stage)
}

func EnsureDir(dirName string) error {

	err := os.Mkdir(dirName, 0755) //probably umasked with 22 not 02

	os.Chmod(dirName, 0755)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func EnsureDirAll(dirName string) error {

	err := os.MkdirAll(dirName, 0755) //probably umasked with 22 not 02

	os.Chmod(dirName, 0755)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func SetupGradexPaths() error {

	paths := []string{
		Root(),
		Ingest(),
		Export(),
		Etc(),
		Var(),
		Usr(),
		Exam(),
		IngestConf(),
		OverlayConf(),
		ExtractConf(),
		SetupConf(),
	}

	for _, path := range paths {

		err := EnsureDirAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetupExamPaths(exam string) error {
	// don't use EnsureDirAll so it flags if we are not otherwise setup
	err := EnsureDir(GetExamPath(exam))
	if err != nil {
		return err
	}

	for _, stage := range ExamStage {
		err := EnsureDir(GetExamStagePath(exam, stage))
		if err != nil {
			return err
		}
	}
	return nil
}
