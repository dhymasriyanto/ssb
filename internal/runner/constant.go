package runner

var (
  opt   *Options
  uhost []string
)

const (
  version = "v0.0.1"
  banner  = `
             _
     ` + version + ` | |
     ___ ___| |__
    / __/ __| '_ \
    \__ \__ \ |_) |
    |___/___/_.__/

Secure Shell Bruteforcer
  infosec@kitabisa.com

`
  usage = `Usage:
  ssb [-p port] [-w wordlist.txt] [-t timeout]
      [-c concurrent] [-o output] [user@]hostname

Options:
  -p port
     Port to connect to on the remote host (default 22).
  -w wordlist
     Path to wordlist file.
  -t timeout
     Connection timeout (default 30s).
  -c concurrent
     Concurrency/threads level (default 100).
  -o output
     Save valid password to file.
  -v
     Verbose mode.

Examples:
  ssb -w wordlist.txt -t 1m -c 1000 root@localhost

`
  timeout = "30s"
)
