#!/usr/bin/env python
import sys
import os
import subprocess
import time
from termcolor import colored

__auth__ = 'jayson.e.grace@gmail.com'

containers = 5

# Open n containers in the default browser
def open_in_browser():
    for n in range(containers):
        os.system("open http://localhost:4440%d" % n)

# Build n vulnerable containers
def build():
    for n in range(containers):
        print(colored("Standing up vulnerablewordpress-%d" % (n), 'green'))
        processes = set()
        processes.add(subprocess.Popen("docker run --name vulnerablewordpress-%d -d -p 4440%d:80 -p 3330%d:3306 wpscanteam/vulnerablewordpress" % (n, n, n), stdout=subprocess.PIPE, shell=True))
        if len(processes) >= n:
            os.wait()
            processes.difference_update([
                p for p in processes if p.poll() is not None])

    time.sleep(15)
    open_in_browser()


# Destroy n vulnerable containers
def destroy():
    for n in range(containers):
        print(colored("Destroying vulnerablewordpress-%d" % (n), 'red'))
        os.system("docker stop vulnerablewordpress-%d" % n)
        os.system("docker rm vulnerablewordpress-%d" % n)

# Validate input to ensure arguments have been input and are valid
def validate_input(args):
    """Validate the number of arguments is correct, and that they're working.
    Obviously, this is not super robust (no further input validation),
    but I'm not writing this to be anything more than a helpful quick and dirty script.
    """
    if len(args) != 1:
        if sys.argv[1] == 'build' or sys.argv[1] == 'destroy':
            return sys.argv[1]
    print("USAGE: python test_lab.py <action>")
    print("EXAMPLE: python test_lab.py build")
    print("EXAMPLE: python test_lab.py destroy")
    sys.exit(-1)

# Execute whatever task has been specified with the input parameters
def execute(task):
    if task == 'build':
        build()
    else:
        destroy()


def main():
    task = validate_input(sys.argv)
    execute(task)


if __name__ == "__main__":
    main()
