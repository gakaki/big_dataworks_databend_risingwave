package csv

import (
    "encoding/csv"
    "os"
)

// WriteToCSV 将 data 写入指定路径的 CSV 文件中
func WriteToCSV(filePath string, data [][]string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, record := range data {
        if err := writer.Write(record); err != nil {
            return err
        }
    }

    return nil
}