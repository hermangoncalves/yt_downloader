package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "youtube-downloader",
	Short: "Um CLI para baixar vídeos ou áudios do YouTube",
	Long:  `youtube-downloader é um CLI em Go que permite baixar vídeos ou áudios do YouTube usando yt-dlp.`,
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}