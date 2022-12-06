package main
import(
    "fmt"
    "os"
    "os/exec"
    "strings"
    "log"
    "text/template"
)

type DayTrait struct {
    Day string
}

const(
    ADVENT_PACKAGE_PREFIX = "github.com/hoduc/AdventOfCode/AdventOfCode2022/"
    UTIL = "util"
    ADVENT_PACKAGE_UTIL = ADVENT_PACKAGE_PREFIX + UTIL
)

var temp = template.Must(template.ParseFiles("day.template"))

func runCommand(command string, args ...string) {
    log.Printf("Running the following command [%v %v]\n", command, strings.Join(args, " "))
    cmd := exec.Command(command, args...)
    // cmd.Dir = dir
    log.Printf("cmd: %v", cmd)
    if err := cmd.Run(); err != nil {
        log.Fatalf("Errors running command: [%v]", err)
        os.Exit(1)
    }
    log.Println("SUCCESS!!")
}

// mkdir day5
// cd day5
// go mod init github.com/hoduc/AdventOfCode/AdventOfCode2022/day5
// go mod edit -replace github.com/hoduc/AdventOfCode/AdventOfCode2022/util=../util
// go mod tidy

func main() {
    if len(os.Args) > 1 {
        fmt.Println("args:", os.Args)
        d := os.Args[1]
        day := "day" + d


        log.Printf("Creating directory: [%v]\n", day);

        if err := os.Mkdir(day, os.ModePerm); err != nil {
            log.Fatalf("Error: [%v]", err)
            os.Exit(1)
        }
        log.Println("SUCCESS!!")


        // write day file
        dayTrait :=  DayTrait{Day: d}
        dayGo := day + "/" + day + ".go"
        log.Printf("Creating %v\n", dayGo)
        dayGoFile, err := os.Create(dayGo)
        if err != nil {
            log.Fatalf("Error crating file: [%v]", err)
            os.Exit(1)
        }

        if err := temp.Execute(dayGoFile, dayTrait); err != nil {
            log.Fatalf("failed template [%v]", err)
            os.Exit(1)
        }

        defer dayGoFile.Close()
        log.Println("SUCCESS!!")



        // cd
        log.Printf("cd %v\n", day)
        if err := os.Chdir(day); err != nil {
            log.Fatalf("Chdir failed: [%v], err")
            os.Exit(1)
        }
        log.Println("SUCCESS!!")
        dir, err := os.Getwd()
        if err != nil {
            log.Fatalf("Error crating file: [%v]", err)
            os.Exit(1)
        }
        log.Println("cwd:", dir)
        // init
        dayPackage := ADVENT_PACKAGE_PREFIX + day
        runCommand("go", []string{"mod", "init", dayPackage}...)

        // replace util
        runCommand("go", []string{"mod", "edit", "-replace", ADVENT_PACKAGE_UTIL + "=../" + UTIL}...)

        // tidy
        runCommand("go", []string{"mod", "tidy"}...)

        // cd back to previous dir
        // and set gowork project file

        log.Printf("cd ../\n", day)
        if err := os.Chdir("../"); err != nil {
            log.Fatalf("Chdir failed: [%v], err")
            os.Exit(1)
        }
        log.Println("SUCCESS!!")

        goWorkDay := "\nuse ./" + day + "\n"
        log.Printf("updating go.work with [%v]\n", goWorkDay)
        goWorkFile, err := os.OpenFile("go.work", os.O_APPEND | os.O_WRONLY, 0644)
        if err != nil {
            log.Fatalf("Error: [%v]", err)
            os.Exit(1)
        }
        defer goWorkFile.Close()
        if _, err := goWorkFile.WriteString(goWorkDay + "\n"); err != nil {
            log.Fatalf("Error: [%v]", err)
            os.Exit(1)
        }
        log.Printf("SUCCESS!!")
        log.Printf("Finished!!!")

    } else {
        fmt.Println("usage: go run day.go [day]")
    }

}
