if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

export ZSH="/root/.oh-my-zsh"

ZSH_THEME="powerlevel10k/powerlevel10k"

source $ZSH/oh-my-zsh.sh

[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh

unsetopt inc_append_history
unsetopt share_history

export CC=/usr/local/gcc-9.1.0/bin/gcc
export CXX=/usr/local/gcc-9.1.0/bin/g++
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/gcc-9.1.0/lib64
export PATH=$PATH:/usr/local/vim/bin
export PATH=$PATH:/usr/local/zsh/bin

alias vi="/usr/local/vim/bin/vim"

alias proxy="export https_proxy=http://192.168.159.1:7890 http_proxy=http://192.168.159.1:7890 all_proxy=socks5://192.168.159.1:7890; curl google.com"
alias gs="git status"
alias gd="git diff"
alias codelines="git ls-files | xargs cat | wc -l"
