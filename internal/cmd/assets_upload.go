package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zbysir/blog/internal/bblog"
	"github.com/zbysir/blog/internal/pkg/log"
	"github.com/zbysir/blog/internal/pkg/oss/qiniu"
	"os"
)

var AssetsUpload = &cobra.Command{
	Use:   "upload",
	Short: "upload file to OSS",
	//Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := viper.GetString("d")
		b, err := bblog.NewBblog(bblog.Option{
			Fs:      os.DirFS(dir),
			ThemeFs: nil,
		})
		if err != nil {
			return err
		}

		conf, err := b.LoadConfig()
		if err != nil {
			return err
		}
		q := qiniu.NewQiniu(conf.Oss.AccessKey, conf.Oss.SecretKey)
		err = q.Uploader().UploadFs(log.StdLogger, conf.Oss.Bucket, "img", os.DirFS("./fe"))
		if err != nil {
			return err
		}

		//c := viper.GetString("config")
		s := viper.AllSettings()
		//log.Infof("config :%+v", c)
		log.Infof("123 %+v %v", args, s)
		return nil
	},
}

func init() {
	AssetsUpload.PersistentFlags().StringP("abc", "a", "", "abc")
	viper.BindPFlag("abc", AssetsUpload.PersistentFlags().Lookup("abc"))
}
