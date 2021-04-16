"""
Does the following:

1. Inits git if used
2. Deletes dockerfiles if not going to be used
3. Deletes config utils if not needed
"""
from __future__ import print_function
import os
import shutil
from subprocess import Popen

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def remove_file(filename):
    """
    generic remove file from project dir
    """
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        os.remove(fullpath)

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ["git", "init"],
        ["git", "add", "--all"],
        ["git", "commit", "-a", "-m", "Initial Commit."]
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=PROJECT_DIRECTORY)
        git.wait()

def remove_db_files():
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, "db"
    ))
    remove_file(os.path.join("action", "db.go"))

def go():
    GOMOD_COMMANDS = [
        ["go", "mod", "vendor"],
        ["gofmt", "-s", "-w", "."]
    ]

    for command in GOMOD_COMMANDS:
        gomod = Popen(command, cwd=PROJECT_DIRECTORY)
        gomod.wait()


if '{{cookiecutter.use_db}}'.lower() == 'none':
    remove_db_files()

if '{{cookiecutter.use_gin}}'.lower() == 'n':
    remove_file(os.path.join("action", "server.go"))
    remove_file(os.path.join("middleware","jwt.go"))
else:
    pattern = 's?__PATH__?'+PROJECT_DIRECTORY+'?'
    wsfile = os.path.join(PROJECT_DIRECTORY,'{{cookiecutter.app_name}}.code-workspace')
    SED_COMMANDS = [
        ["sed", "-i", pattern, wsfile]
    ]
    for command in SED_COMMANDS:
        sed = Popen(command, cwd=PROJECT_DIRECTORY)
        sed.wait()

# Initialize Go Modules
go()

# Initialize Git (should be run after all file have been modified or deleted)
init_git()