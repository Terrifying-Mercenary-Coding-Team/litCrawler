import subprocess
import json

if __name__ == "__main__":
    urls = ["http://codeforces.com/",
            "https://leetcode.com/", "https://github.com/"]
    dumped_urls = json.dumps(urls)
    proc = subprocess.Popen(
        ["./URLChecker", dumped_urls], stdout=subprocess.PIPE)
    output, err = proc.communicate()
    output = json.loads(output.decode('utf-8'))
    print(output)
