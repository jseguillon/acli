#!/usr/bin/env bash 
set -e
export VERSION=$(curl -s https://api.github.com/repos/jseguillon/chat-gpt-cli/releases | grep tag_name | grep -v -- '-rc' | sort -r | head -1 | awk -F': ' '{print $2}' | sed 's/,//' | xargs)

shell_install_func() {
    if [ -f $1 ]; then
        set +e; ask "Add aliases and functions in $1"; ret="$?"; set -e
        if [ $ret -eq 1 ]; then
            echo "$INSTALL_SH_FUNC" >> $1
        fi
    fi
}

shell_install_config() {
    if [ -f $1 ]; then
        set +e; ask "Add default configuration in $1"; ret="$?"; set -e
        if [ $ret -eq 1 ]; then
            echo "$INSTALL_SH_KEY" >> $1
        fi
    fi
}

shell_install() {
echo
echo "chat-gpt-cli aliases and functions:"
echo "\`\`\`"
echo "$INSTALL_SH_FUNC"
echo "\`\`\`"
echo

files=($HOME/.bashrc $HOME/.zshrc)
for file in ${files[@]}; do
    shell_install_func $file
done

echo 
echo 
echo "Default chat-gpt-cli configuration:"
echo "\`\`\`"
echo "$INSTALL_SH_KEY"
echo "\`\`\`"
echo

files=($HOME/.bashrc $HOME/.zshrc)
for file in ${files[@]}; do
    shell_install_config $file 
done
}

ask() {
    t="$(echo -e '\t')"
    read -p "$1 ?$t[Y/n] " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Nn]$ ]]
    then
        return 0
    else
        return 1
    fi
}

architecture=""
case $(uname -m) in
    i386)    architecture="386" ;;
    i686)    architecture="386" ;;
    x86_64)  architecture="amd64" ;;
    aarch64) architecture="arm64" ;;
    arm64)   architecture="arm64" ;;
    arm)     dpkg --print-architecture | grep -q "arm64" && architecture="arm64" || architecture="arm" ;;
    *)       echo "Architecture is not supported"; exit -1 ;;
esac

os=""
# Test system OS
case $(uname -s) in
    Linux)
        os="linux" ;;
    NetBSD)
        os="netbsd" ;;
    NetBSD)
        os="netbsd" ;;
    FreeBSD)
        os="freebsd" ;;
    Darwin)
        os="darwin" ;;
    *)
        echo "System OS is not supported"; exit -1 ;;
esac

install_dir="/usr/local/bin"
echo "# Install"
echo -n "Destination directory ? "
read -p "($install_dir) : " -r
if [ ! $REPLY == "" ]; then install_dir="$REPLY"; fi
echo $install_dir

echo "Dowloading $url into $install_dir/chat-gpt-cli"
url="https://github.com/jseguillon/chat-gpt-cli/releases/download/$VERSION/chat-gpt-cli-$os-$architecture"
dest="$install_dir/chat-gpt-cli"
sudo curl -sSL $url -o $dest
sudo chmod +x $dest

set +e
read -r -d '' INSTALL_SH_KEY <<'EOF'
CHAT_GPT_API_KEY="[Get your key at: https://beta.openai.com/account/api-keys]"
EOF
read -r -d '' INSTALL_SH_FUNC <<'EOF'
alias fix='eval $(chat-gpt-cli --script fixCmd "$(fc -nl -1)" $?)'
howto() { h="$@"; eval $(chat-gpt-cli --script howCmd "$h") ; }
EOF
set -e

echo
echo "# Configure"
shell_install

echo "Installation done."
echo 
echo

echo "# Sample usage:"
echo
echo "* use 'chat-gpt-cli' for discussions or complex task solving: "
echo '     > chat-gpt-cli "can GPT help me for daily command line tasks ?"'
echo '     > chat-gpt-cli "[very complex script description for bash/javascript/python/etc...]"'
echo
echo "* use 'howto' function for quick one liner answers and interactive mode: "
echo '    > howto openssl test SSL expiracy of github.com'
echo '    > howto "find all files more than 30Mb "'
echo
echo "* use 'fix' for quick fixing typos: "
echo "    [run typo command like 'rrm', 'lls', 'cd..', etc..]"
echo "    then type 'fix' and get fixed command ready to run"
echo 

echo "If you like, please give a star on github.com or follow @Jseguillon on Twitter. Thanks."
echo
