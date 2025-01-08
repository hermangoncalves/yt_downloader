package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// Flags
var url string
var format string
var output string

// downloadCmd implementa o subcomando `download`
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Baixa vídeos ou áudios do YouTube",
	Long:  `Permite baixar vídeos ou áudios do YouTube, escolhendo o formato (vídeo ou áudio).`,
	Run: func(cmd *cobra.Command, args []string) {
		// Validação da URL
		if url == "" {
			fmt.Println("Erro: A flag --url é obrigatória.")
			os.Exit(1)
		}

		// Determina o formato
		var formatOption string
		switch format {
		case "video":
			formatOption = "best"
		case "audio":
			formatOption = "bestaudio"
		default:
			fmt.Println("Erro: Formato inválido. Use 'video' ou 'audio'.")
			os.Exit(1)
		}

		// Monta o comando yt-dlp
		args = []string{"-f", formatOption, url}

		if output != "" {
			args = append(args, "-o", output)
		}

		// Executa o comando yt-dlp
		cmdExec := exec.Command("yt-dlp", args...)
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr

		fmt.Println("Baixando...")
		if err := cmdExec.Run(); err != nil {
			fmt.Printf("Erro ao executar yt-dlp: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Download concluído com sucesso!")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&url, "url", "u", "", "URL do vídeo do YouTube (obrigatório)")
	downloadCmd.Flags().StringVarP(&format, "format", "f", "video", "Formato do download: 'video' ou 'audio'")
	downloadCmd.Flags().StringVarP(&output, "output", "o", "", "Nome do arquivo de saída (opcional)")
}
