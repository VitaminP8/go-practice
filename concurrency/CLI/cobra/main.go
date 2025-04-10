package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Определение глобальных флагов (общих для всех команд)
var (
	verbose bool
)

// Корневая команда
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "Демонстрационное приложение с использованием Cobra",
	Long:  "Это демонстрационное приложение для демонстрации использования библиотеки Cobra на Go.",
	// Run - функция, которая будет вызвана при выполнении команды
	Run: func(cmd *cobra.Command, args []string) {
		// Используйте глобальные флаги здесь, если нужно
		if verbose {
			fmt.Println("Запуск в режиме подробного вывода")
		}
		fmt.Println("Привет из корневой команды!")
	},
}

// Команда 1
var cmd1 = &cobra.Command{
	Use:   "command1",
	Short: "Описание команды 1",
	Long:  "Это описание для команды 1",
	// Run - функция, которая будет вызвана при выполнении этой команды
	Run: func(cmd *cobra.Command, args []string) {
		// Доступ к глобальным флагам
		if verbose {
			fmt.Println("Запуск command1 в режиме подробного вывода")
		}
		fmt.Println("Привет из команды 1!")
	},
}

// Команда H
var cmdH = &cobra.Command{
	Use:   "help",
	Short: "Описание команды 2",
	Long:  "Это описание для команды 2",
	// Run - функция, которая будет вызвана при выполнении этой команды
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Привет из команды 2!")
	},
}

// Флаги, специфичные для Command 1
var cmd1Flag string

func main() {
	// Добавление глобальных флагов к корневой команде
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Включить подробный режим")
	rootCmd.PersistentFlags().BoolVar(&verbose, "other_long_param", false, "Пример флага без поддержки сокращений")

	// Добавление флагов, специфичных для Command 1
	_ = rootCmd.Flags().BoolP("help", "h", false, "hackerhelp")
	cmd1.Flags().StringVarP(&cmd1Flag, "flag1", "f", "", "Флаг, специфичный для Command 1")

	// Добавление Command 1 к корневой команде
	rootCmd.AddCommand(cmd1)
	rootCmd.AddCommand(cmdH)

	// Выполнение приложения Cobra
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
