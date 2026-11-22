package config

func BackupProcess(bFolder string) error {
	repo, err := ReadConf()
	if err != nil {
		return err
	}

	repo.Source = bFolder
	if err = WriteConf(repo); err != nil {
		return err
	}

	return nil
}
