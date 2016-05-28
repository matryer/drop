package main

func gimme(source, path, dest string) ([]string, error) {
	repo, done, err := goget(source, path)
	if err != nil {
		return nil, err
	}
	defer done()
	files, err := files(repo.path)
	if err != nil {
		return nil, err
	}

	license, err := license(repo.path)
	if err != nil {
		return nil, err
	}

	copiedFiles, err := copy(dest, source, path, repo.id, license, files...)
	if err != nil {
		return nil, err
	}
	return copiedFiles, nil
}
