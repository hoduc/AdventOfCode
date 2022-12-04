package util

import (
    "os"
    "log"
    "bufio"
    "strings"
)

type onlineFn func(string) error

func ReadLines(fileName string, onLine onlineFn) error {
    f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatalf("Error open file: %v", err)
        return err
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        if err := onLine(line); err != nil {
            log.Fatalf("onLine failed on [%v]", err)
            return err
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Scan file err: %v", err)
        return err
    }

    return nil
}

func ReadLinesEmbed(lines string, onLine onlineFn) error {
    scanner := bufio.NewScanner(strings.NewReader(lines))
    for scanner.Scan() {
        line := scanner.Text()
        if err := onLine(line); err != nil {
            log.Fatalf("onLine failed on [%v]", err)
            return err
        }
    }
    return nil
}
