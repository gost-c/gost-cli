package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gost-c/gost-cli/colors"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	// BaseURL is the base url of gost api services
	BaseURL = "https://iuao0sjxmi.execute-api.ap-southeast-1.amazonaws.com/development/"
	//BaseURL = "http://localhost:9393/"

	// WebURL is the base url of gost web services
	WebURL = "https://gost.zcong.moe/#/gost/"
	space  = "  "
	// MaxSize is single file max size, 10k
	MaxSize = int64(1024 * 10)
	// MaxSizeHuman is max size for human
	MaxSizeHuman = humanize.Bytes(uint64(MaxSize))
	// MaxFilesCount is max files count allowed
	MaxFilesCount = 10
	// ErrMaxFilesCount is error message for too many files
	ErrMaxFilesCount = fmt.Errorf("More than %d files is not allowed ", MaxFilesCount)
)

// PathStat is file stat
type PathStat struct {
	// Error is error message
	Error error
	// IsFolder show if the PathStat is folder
	IsFolder bool
	// Size is file humanized size if is a file
	Size string
}

// ConfigFile is the true path of config file `.gostrc`
var ConfigFile string

const config = ".gostrc"

func init() {
	home, err := homedir.Dir()
	if err != nil {
		Fail("An error occurred when get user home!")
		os.Exit(1)
	}
	ConfigFile = path.Join(home, config)
}

// WriteConfig write token to config file
func WriteConfig(token []byte) error {
	return ioutil.WriteFile(ConfigFile, token, 0644)
}

// ReadConfig read token from config file
func ReadConfig() ([]byte, error) {
	return ioutil.ReadFile(ConfigFile)
}

// Success log success message with colors
func Success(str string) {
	fmt.Println()
	fmt.Printf("%s%s%s", colors.Green("SUCCESS"), space, str)
	fmt.Println()
}

// Fail log error message with colors
func Fail(str string) {
	fmt.Println()
	fmt.Printf("%s%s%s", colors.Red("ERROR"), space, str)
	fmt.Println()
}

// DoRequest do a http request and decode the json response
func DoRequest(method, url string, v interface{}, headers map[string]string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(""))
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

// GetPathStat return the path type, if is file return file size
func GetPathStat(p string) *PathStat {
	fileInfo, err := os.Stat(p)
	ps := &PathStat{}
	if err != nil {
		ps.Error = err
		return ps
	}
	if fileInfo.IsDir() {
		ps.IsFolder = true
		return ps
	}
	s := fileInfo.Size()
	if s > MaxSize {
		ps.Error = errFileTooBig(s)
		return ps
	}
	ps.IsFolder = false
	ps.Size = humanize.Bytes(uint64(fileInfo.Size()))
	return ps
}

func errFileTooBig(s int64) error {
	return fmt.Errorf("File is too big %s, max allowed size is %s ", humanize.Bytes(uint64(s)), MaxSizeHuman)
}

// Pad helper.
func Pad() func() {
	println()
	return func() {
		println()
	}
}

// LogPad outputs a log message with padding.
func LogPad(msg string) {
	defer Pad()()
	fmt.Println(msg)
}
