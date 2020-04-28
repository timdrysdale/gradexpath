package gradexpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	isTesting bool
	testroot  = "./tmp-delete-me"
	ExamStage = []string{
		config,
		acceptedPapers,
		acceptedReceipts,
		tempImages,
		tempPages,
		anonPapers,
		markerReady,
		markerSent,
		markerBack,
		moderatorReady,
		moderatorSent,
		moderatorBack,

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

const (
	config           = "00-config"
	acceptedReceipts = "02-accepted-receipts"
	acceptedPapers   = "03-accepted-papers"
	tempImages       = "03-temporary-images"
	tempPages        = "04-temporary-pages"
	anonPapers       = "05-anonymous-papers"
	markerReady      = "20-marker-ready"
	markerSent       = "21-marker-sent"
	markerBack       = "22-marker-back"
	moderatorReady   = "30-moderator-ready"
	moderatorSent    = "31-moderator-sent"
	moderatorBack    = "32-moderator-back"
	checkerReady     = "30-checker-ready"
	checkerSent      = "31-checker-sent"
	checkerBack      = "32-checker-back"

	N = 3
)

// Path is in the package name, so anything without another noun IS a path
// non path things need a noun

func limit(initials string, N int) string {
	if len(initials) < 3 {
		N = len(initials)
	}
	return strings.ToUpper(initials[0:N])
}

func DoneDecoration() string {
	return "d"
}

func MarkerABCDecoration(initials string) string {
	return fmt.Sprintf("-ma%s", limit(initials, N))
}

func MarkerABCDirName(initials string) string {
	return limit(initials, N)
}

func ModeratorABCDecoration(initials string) string {
	return fmt.Sprintf("-mo%s", limit(initials, N))
}

func ModeratorABCDirName(initials string) string {
	return limit(initials, N)
}

func CheckerABCDecoration(initials string) string {
	return fmt.Sprintf("-c%s", limit(initials, N))
}

func CheckerABCDirName(initials string) string {
	return limit(initials, N)
}

func MarkerNDecoration(number int) string {
	return fmt.Sprintf("-ma%d", number)
}

func MarkerNDirName(number int) string {
	return fmt.Sprintf("marker%d", number)
}

func ModeratorNDecoration(number int) string {
	return fmt.Sprintf("-mo%d", number)
}

func ModeratorNDirName(number int) string {
	return fmt.Sprintf("moderator%d", number)
}

func CheckerNDecoration(number int) string {
	return fmt.Sprintf("-c%d", number)
}

func CheckerNDirName(number int) string {
	return fmt.Sprintf("checker%d", number)
}

func MarkerReady(exam, marker string) string {
	return filepath.Join(Exam(), exam, markerReady, marker)
}

func MarkerSent(exam, marker string) string {
	return filepath.Join(Exam(), exam, markerSent, marker)
}

func MarkerBack(exam, marker string) string {
	return filepath.Join(Exam(), exam, markerBack, marker)
}

func ModeratorReady(exam, moderator string) string {
	return filepath.Join(Exam(), exam, moderatorReady, moderator)
}

func ModeratorSent(exam, moderator string) string {
	return filepath.Join(Exam(), exam, moderatorSent, moderator)
}

func ModeratorBack(exam, moderator string) string {
	return filepath.Join(Exam(), exam, moderatorBack, moderator)
}

func CheckerReady(exam, checker string) string {
	return filepath.Join(Exam(), exam, checkerReady, checker)
}

func CheckerSent(exam, checker string) string {
	return filepath.Join(Exam(), exam, checkerSent, checker)
}

func CheckerBack(exam, checker string) string {
	return filepath.Join(Exam(), exam, checkerBack, checker)
}

func FlattenLayoutSVG() string {
	return filepath.Join(IngestTemplate(), "layout-flatten-312pt.svg")
}

func OverlayLayoutSVG() string {
	return filepath.Join(OverlayTemplate(), "layout.svg")
}

func AcceptedPapers(exam string) string {
	return filepath.Join(Exam(), exam, acceptedPapers)
}

func AcceptedReceipts(exam string) string {
	return filepath.Join(Exam(), exam, acceptedReceipts)
}

//TODO in flatten, swap these paths for the general named ones below
func AcceptedPaperImages(exam string) string {
	return filepath.Join(Exam(), exam, tempImages)
}
func AcceptedPaperPages(exam string) string {
	return filepath.Join(Exam(), exam, tempPages)
}
func PaperImages(exam string) string {
	return filepath.Join(Exam(), exam, tempImages)
}
func PaperPages(exam string) string {
	return filepath.Join(Exam(), exam, tempPages)
}

func AnonymousPapers(exam string) string {
	return filepath.Join(Exam(), exam, anonPapers)
}

func Identity() string {
	return filepath.Join(Etc(), "identity")
}

func IdentityCSV() string {
	return filepath.Join(Identity(), "identity.csv")
}

func Ingest() string {
	return filepath.Join(Root(), "ingest")
}

func IngestTemplate() string {
	return filepath.Join(IngestConf(), "template")
}

func OverlayTemplate() string {
	return filepath.Join(OverlayConf(), "template")

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
		Identity(),
		Export(),
		Var(),
		Usr(),
		Exam(),
		TempPdf(),
		TempTxt(),
		Etc(),
		IngestConf(),
		OverlayConf(),
		IngestTemplate(),
		OverlayTemplate(),
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
func MoveIfNewerThanDestination(source, destination string) error {

	//check both exist
	sourceInfo, err := os.Stat(source)

	if err != nil {
		return err
	}

	destinationInfo, err := os.Stat(destination)

	// source newer by definition if destination does not exist
	if os.IsNotExist(err) {
		err = os.Rename(source, destination)
		return err
	}
	if err != nil {
		return err
	}
	if sourceInfo.ModTime().After(destinationInfo.ModTime()) {
		err = os.Rename(source, destination)
		return err
	}

	return nil

}

func MoveIfNewerThanDestinationInDir(source, destinationDir string) error {

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
	if err != nil {
		return err
	}
	if sourceInfo.ModTime().After(destinationInfo.ModTime()) {
		err = os.Rename(source, destination)
		return err
	}

	return nil

}

func ExamDiet(exam string) string {

	m := int(time.Now().Month())

	switch {
	case m > 4 && m < 6:
		return fmt.Sprintf("May-%d", time.Now().Year())
	case m > 6 && m < 10:
		return fmt.Sprintf("Aug-%d", time.Now().Year())
	case m > 10 || m < 3:
		return fmt.Sprintf("Dec-%d", time.Now().Year())
	default:
		return fmt.Sprintf("%d", time.Now().Year())
	}

	return ""
}
