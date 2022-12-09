package util

import (
    "os"
    "log"
    "bufio"
    "strings"
    "math"
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

// TODO: More type
func Max(a, b interface{}) interface{} {
    switch a.(type) {
    case int:
        return int(MaxInt(int64(a.(int)), int64(b.(int))))
    case int32:
        return int32(MaxInt(int64(a.(int32)), int64(b.(int32))))
    case int64:
        return MaxInt(a.(int64), b.(int64))
    case float32:
        return float32(math.Max(float64(a.(float32)), float64(b.(float32))))
    case float64:
        return math.Max(a.(float64), b.(float64))
    }
    return nil
}

func MaxInt(a, b int64) int64 {
    if a < b {
        return b
    }
    return a
}

func Abs(x interface{}) interface {}{
    switch x.(type) {
    case int:
        return int(AbsInt(int64(x.(int))))
    case int32:
        return int32(AbsInt(int64(x.(int32))))
    case int64:
        return AbsInt(x.(int64))
    case float32:
        return float32(math.Abs(float64(x.(float32))))
    case float64:
        return math.Abs(x.(float64))
    }
    return nil
}

func AbsInt(x int64) int64 {
    if (x < 0) {
        return -x
    }
    return x
}
