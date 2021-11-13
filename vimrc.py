"" 基础选项
let mapleader = ","                       " Rebind <Leader> key
syntax enable                             " 语法高亮
filetype plugin indent on                 " 开启文件类型检测（必须）
set nocompatible                          " 关闭 vi 兼容模式
set backspace=indent,eol,start            " 在 insert mode 下允许使用 backspace 键
set history=1000                          " 指定 history 行数
set novisualbell                          " No sounds
set noerrorbells                          " No noise
set confirm                               " 存在 .swp 文件时需要确认
set nofoldenable                          " 不启用折叠
set hlsearch                              " 高亮搜索关键字
set ignorecase                            " 搜索/替换时大小写敏感

" 编码选项
set fileencodings=utf-8,gb2312,gbk
set fileencoding=utf-8
set termencoding=utf-8
set encoding=utf-8
set fileformat=unix

" 布局选项
set number                                " 显示行号
set numberwidth=4                         " 指定行号宽度
set textwidth=79                          " 指定 text 宽度，PEP8 为 79
set colorcolumn=79                        " 指定宽度高亮
set splitright                            " 指定窗口分割时把新窗口放到当前窗口之右
set list listchars=tab:>.,trail:.         " 显示 tab、空格的制表符

" 缩进选项
set tabstop=4                             " 指定 tap 宽度
set expandtab                             " 用空格代替 tab
set shiftwidth=4                          " 指定制表符的宽度
set softtabstop=4                         " 指定自动缩进的宽度
set showmatch                             " 显示匹配的括号
set autoindent
set smartindent

" 在 insert mode 下开启光标线和光标柱，普通模式时关闭
autocmd InsertLeave * call ToggleCursor()
autocmd InsertEnter * call ToggleCursor()
function ToggleCursor() abort
    set cursorline!
    set cursorcolumn!
endfunction

"""""""""""""""""""""""""""""""""""""""" Vundle Plugins
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
Plugin 'VundleVim/Vundle.vim'
Plugin 'altercation/vim-colors-solarized'
Plugin 'altercation/solarized'
Plugin 'powerline/powerline', {'rtp': 'powerline/bindings/vim'}
Plugin 'scrooloose/nerdtree'
Plugin 'jistr/vim-nerdtree-tabs'
Plugin 'Xuyuanp/nerdtree-git-plugin'
Plugin 'preservim/tagbar'
Plugin 'Valloric/YouCompleteMe'
Plugin 'python-mode/python-mode'
Plugin 'easymotion/vim-easymotion'
call vundle#end()


" solarized
colorscheme solarized
set t_Co=256
set background=dark
let g:solarized_termcolors=16
let g:solarized_visibility='high'
let g:solarized_contrast='high'


" powerline
set laststatus=2                          " Always show status line
set statusline+=%F


" nerdtree
let NERDChristmasTree=0
let NERDTreeWinSize=35
let NERDTreeChDirMode=2
let NERDTreeIgnore=['\~$', '\.pyc$', '\.swp$', '^__pycache__$']
let NERDTreeShowBookmarks=1
let NERDTreeWinPos="left"
autocmd vimenter * if !argc() | NERDTree | endif
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTreeType") && b:NERDTreeType == "primary") | q | endif
nmap <leader>n :NERDTreeToggle<CR>


" tagbar
let g:tagbar_width=35
let g:tagbar_autofocus=1
let g:tagbar_ctags_bin ='/usr/local/bin/ctags'
nmap <leader>t :TagbarToggle<CR>


" YouCompleteMe
let g:ycm_global_ycm_extra_conf='/root/.vim/bundle/YouCompleteMe/.ycm_extra_conf.py'
let g:ycm_min_num_of_chars_for_completion=2                 " 两个 chars 后开始补全
let g:ycm_autoclose_preview_window_after_completion=1       " auto-close preview window after select a completion string
let g:ycm_autoclose_preview_window_after_insertion=1        " auto-close preview window after leaving insert mode
let g:ycm_confirm_extra_conf = 0                            " 关闭额外配置
let g:ycm_complete_in_comments = 1                          " 在 comment 中补全
let g:ycm_complete_in_strings = 1                           " 在 string 中补全
let g:ycm_seed_identifiers_with_syntax=1                    " 从 syntax 中收集 identifiers
let g:ycm_collect_identifiers_from_tags_files=1             " 从 tag files 中收集 identifiers
let g:ycm_collect_identifiers_from_comments_and_strings = 1 " 从 comment 和 string 中收集 identifiers
let g:syntastic_always_populate_loc_list = 1
let g:ycm_cache_omnifunc=0
let g:ycm_error_symbol = '>>'
let g:ycm_warning_symbol = '>*'
let g:ycm_key_invoke_completion='<C-/>'


" python-mode
let g:pymode_options = 1
let g:pymode_python = 'python3'
let g:pymode_rope_completion = 0          " 关闭 python-mode 的自动补全，使用 YCM
let g:pymode_indent = 1                   " 使用 PEP8 风格的缩进
let g:pymode_syntax = 1                   " 使用 Python 语法高亮
let g:pymode_syntax_all = 1
let g:pymode_syntax_indent_errors = g:pymode_syntax_all
let g:pymode_syntax_space_errors = g:pymode_syntax_all
let g:pymode_breakpoint = 1               " 使用 python-mode 设置断点
let g:pymode_breakpoint_bind = '<leader>b'
let g:pymode_rope_goto_definition_bind = '<C-s>' " 显示 function/method defines
let g:pymode_lint = 1                     " 使用 lint 语法检查
let g:pymode_lint_on_write = 1
let g:pymode_lint_on_fly = 0              " 关闭编辑时进行检查
let g:pymode_lint_signs = 1               " 显示 python-mode 相关的标志
let g:pymode_lint_todo_symbol = 'WW'
let g:pymode_lint_comment_symbol = 'CC'
let g:pymode_lint_visual_symbol = 'RR'
let g:pymode_lint_error_symbol = 'EE'
let g:pymode_lint_info_symbol = 'II'
let g:pymode_lint_pyflakes_symbol = 'FF'
let g:pymode_lint_checkers = ['pylint', 'pep8']
let g:pymode_virtualenv = 1               " 使用 virtualenv
let g:pymode_motion = 1                   " 使用快捷移动
let g:pymode_trim_whitespaces = 1         " 保存文件时自动删除无用空格
let g:pymode_warnings = 0                 " 开启警告


" easymotion
let g:EasyMotion_smartcase = 1
map <Leader><leader>h <Plug>(easymotion-linebackward)
map <Leader><Leader>k <Plug>(easymotion-j)
map <Leader><Leader>j <Plug>(easymotion-k)
map <Leader><leader>l <Plug>(easymotion-lineforward)
map <Leader><leader>. <Plug>(easymotion-repeat)
