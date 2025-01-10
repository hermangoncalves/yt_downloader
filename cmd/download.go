package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	tempDir        = "downloads"
	formatVideo    = "video"
	formatAudio    = "audio"
	audioFormatMP3 = "mp3" 
)

// Flags
var url string
var format string
var output string


var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Baixa vídeos ou áudios do YouTube",
	Long:  `Permite baixar vídeos ou áudios do YouTube, escolhendo o formato (vídeo ou áudio).`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runDownload(); err != nil {
			log.Fatalf("Erro: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Define flags for the download command
	downloadCmd.Flags().StringVarP(&url, "url", "u", "", "URL do vídeo do YouTube (obrigatório)")
	downloadCmd.Flags().StringVarP(&format, "format", "f", formatVideo, "Formato do download: 'video' ou 'audio'")
	downloadCmd.Flags().StringVarP(&output, "output", "o", "", "Nome do arquivo de saída (opcional)")
}


func runDownload() error {
	log.Println("Iniciando processo de download...")

	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return fmt.Errorf("falha ao criar diretório temporário: %v", err)
	}

	if url == "" {
		return fmt.Errorf("a flag --url é obrigatória")
	}

	formatOption, err := getFormatOption(format)
	if err != nil {
		return err
	}


	downloadArgs := buildDownloadArgs(formatOption, url, output)


	if err := executeDownloadCommand(downloadArgs); err != nil {
		return fmt.Errorf("falha ao executar o comando yt-dlp: %v", err)
	}

	log.Println("Download concluído com sucesso!")
	return nil
}

func getFormatOption(format string) (string, error) {
	switch format {
	case formatVideo:
		return "best", nil
	case formatAudio:
		return "bestaudio", nil
	default:
		return "", fmt.Errorf("formato inválido: use 'video' ou 'audio'")
	}
}

func buildDownloadArgs(formatOption, url, output string) []string {
	args := []string{"-f", formatOption, url}

	if format == formatAudio {
		args = append(args, "--extract-audio", "--audio-format", audioFormatMP3)
	}

	if output != "" {
		args = append(args, "-o", filepath.Join(tempDir, output))
	} else {
		args = append(args, "-o", filepath.Join(tempDir, "%(title)s.%(ext)s"))
	}

	return args
}

func executeDownloadCommand(args []string) error {
	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Baixando...")
	return cmd.Run()
}
