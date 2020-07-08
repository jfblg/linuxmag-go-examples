package main

func driveWatch(done chan error) chan string {
	seen := map[string]bool{}
	init := true
	drivech := make(chan string)

	go func() {
		for {
			dpaths, err := devices()
			if err != nil {
				done <- err
			}
			for _, dpath := range dpaths {
				if _, ok := seen[dpaht]; !ok {
					seen[dpath] = true
					if !init {
						drivech <- dpath
					}
				}
			}
			init = false
			time.Sleep(1 * time.Second)
		}
	}()
	return drivech
}

func driveSize(path string) (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("sfdisk", "-s", path)
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	sizeStr := strings.TrimSuffix(out.String(), "\n")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.1f GB", float64(size)/float64(1024*1024)), nil
}

func devices() ([]string, error) {
	devices := []string{}
	paths, _ := filepath.Glob("/dev/sd*")
	if len(paths) == 0 {
		return devices, errors.New("No devices found")
	}
	for _, path := range paths {
		devices = append(devices, filepath.Base(path))
	}
	return devices, nil
}