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
		"10-temp-pdf",
		"11-temp-txt",
		"12-temp-annotate",
		"13-temp-reject",
		"14-temp-paper",
		"15-temp-mark",
		"16-temp-moderate",
		"17-temp-check",
		"18-temp-config",
		"19-temp-ignore",
		"20-complete-paper-set",
		"21-complete-receipt-set",
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

func TempPdf() string {
	return filepath.Join(Root(), "temp-pdf")
}

func TempTxt() string {
	return filepath.Join(Root(), "temp-txt")
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

func SetupGradexPaths() error {

	paths := []string{
		Root(),
		Ingest(),
		Export(),
		Var(),
		Usr(),
		Exam(),
		TempPdf(),
		TempTxt(),
		Etc(),
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

// if the source file is not newer, it's not an error
// we just won't move it - anything left we deal with later
func MoveIfNewerThanDestination(source, destinationDir string) error {

	//check both exist
	sourceInfo, err := os.Stat(source)

	if err != nil {
		return err
	}

	destination := filepath.Join(destinationDir, filepath.Base(source))

	destinationInfo, err := os.Stat(destination)

	// source newer by definition if destination does not exist
	if os.IsNotExist(err) {
		err = os.Rename(source, destination)
		return err
	}

	if sourceInfo.ModTime().After(destinationInfo.ModTime()) {
		err = os.Rename(source, destination)
		return err
	}

	return nil

}
