import subprocess

if __name__ == "__main__":
    proc = subprocess.Popen(
        ["./RandomPicker"], stdout=subprocess.PIPE)
    output, err = proc.communicate()
    print(output.decode('utf-8'))
