" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" Vundle
Plugin 'VundleVim/Vundle.vim'

" Solarized
Plugin 'altercation/vim-colors-solarized'

" Nerdtree
Plugin 'scrooloose/nerdtree'
Plugin 'jistr/vim-nerdtree-tabs'

" Tagbar
Plugin 'majutsushi/tagbar'

" YouCompleteMe
Plugin 'Valloric/YouCompleteMe'

" Syntastic
Plugin 'scrooloose/syntastic'

" vim-easymotion
Plugin 'easymotion/vim-easymotion'

" Powerline（卡顿）
Plugin 'Lokaltog/powerline', {'rtp': 'powerline/bindings/vim/'}
Plugin 'powerline/fonts'

" vim-go
Plugin 'fatih/vim-go'

call vundle#end()            " required
filetype plugin indent on    " required
filetype plugin on


"=============================================== General Config =================================================
set nocompatible                          " required
filetype off                              " required
set number                                " 显示行号
set ruler                                 " 打开状态栏标尺
set backspace=indent,eol,start            " 插入模式下允许退格
set fileencodings=utf-8,gbk               " 设置文件编码
set laststatus=2                          " 显示状态栏
set history=1000                          " history 数量
set showcmd                               " 在底部显示指令
set showmode                              " 在底部显示当前的模式
set showmatch                             " 输入 )/} 时，光标会暂时的回到相匹配的 (/{
set gcr=a:blinkon0                        " 禁用光标闪烁
set novisualbell                          " 没有声音
set noerrorbells                          " 没有噪音
set autoread                              " 重新加载在 vim 外部更改的文件
set statusline+=%{fugitive#statusline()}  " Git Hotness
"set list listchars=tab:>.,trail:.         " 直观地显示制表符和尾随空格
set nolist
set linebreak                             " 在方便的地方换行
set wildmenu                              " 命令自动完成
set shortmess=atI
set statusline=\ %<%F[%1*%M%*%n%R%H]%=\ %y\ %0(%{&fileformat}\ %{&encoding}\ %c:%l/%L%)\
set nobackup
set confirm
set nowb
set textwidth=79                          " 显示 79 个字符长度的位置
set colorcolumn=79
highlight ColorColumn ctermbg=gray
set numberwidth=4
set fileformat=unix
set expandtab
set cmdheight=1
set ignorecase smartcase

" 突出当前行和列（卡顿）
au WinLeave * set nocursorline nocursorcolumn
au WinEnter * set cursorline cursorcolumn
set cursorline cursorcolumn
set clipboard=unnamed

" split navigations
set splitbelow
set splitright
nnoremap <C-J> <C-W><C-J>
nnoremap <C-K> <C-W><C-K>
nnoremap <C-L> <C-W><C-L>
nnoremap <C-H> <C-W><C-H>

" Persistent Undo
set undodir=~/.vim/backups
set undofile

" 搜索
set incsearch                   " 输入搜索内容时就显示搜索结果
set hlsearch                    " 搜索时高亮显示被找到的文本
set viminfo='100,f1             " Save up to 100 marks, enable capital marks

" 缩进
set autoindent
set smartindent                 " 开启新行时使用智能自动缩进
set smarttab
set shiftwidth=4                " 设定 << 和 >> 命令移动时的宽度为 4
set softtabstop=4               " 使得按退格键时可以一次删掉 4 个空格
set tabstop=4                   " 设定 Tab 长度为 4
set expandtab

" Folds
set foldenable
set foldcolumn=0
set foldnestmax=3               " Deepest fold is 3 levels
set nofoldenable                " Dont fold by default

set foldmethod=indent           " Fold based on indent
set foldlevel=99
" use space to open folds
nnoremap <space> za

"set foldclose=all
nnoremap <space> @=((foldclosed(line('.')) < 0) ? 'zc' : 'zo')<CR>

" Leader setting
let mapleader = ","             " Rebind <Leader> key

" Syntax Highlight
syntax on

" Run commands that require an interactive shell
nnoremap <Leader>r :RunInInteractiveShell<space>


"=============================================== Solarized 主题
syntax enable
set background=dark                " 使用阴主题
let g:solarized_visibility='high'  " 设置空白符的可见性
let g:solarized_contrast='high'    " 设置对比度
try
  colorscheme solarized            " 设定配色方案
catch /^Vim\%((\a\+)\)\=:E185/
endtry


"=============================================== NERD Tree 目录树窗口
" autocmd vimenter * NERDTree          " 自动开启 Nerdtree
let g:NERDTreeHidden=0               "不显示隐藏文件
let NERDTreeWinSize=35               " 设定 NERDTree 视窗大小
let NERDTreeChDirMode=2              " 根据 NERDTree 中的根目录更改当前工作目录
let NERDTreeIgnore=['\~$', '\.pyc$', '\.swp$']    " 忽略清单
let NERDTreeShowBookmarks=1          " 开启 Nerdtree 时自动显示 Bookmarks
let NERDTreeWinPos="left"            " 窗口位置
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTreeType") && b:NERDTreeType == "primary") | q | endif  " 当 NERDTree 为剩下的唯一窗口时自动关闭
nmap <leader>d :NERDTreeToggle<CR>


"=============================================== Tagbar 符号窗口
let g:tagbar_width=35       " 窗口宽度
let g:tagbar_right_=1       " 窗口位置
let g:tagbar_autofocus=1    " 窗口打开光标自动进入
let g:tagbar_ctags_bin ='ctags'
nmap <leader>c :TagbarToggle<CR>

let g:tagbar_type_go = {
    \ 'ctagstype' : 'go',
    \ 'kinds'     : [
        \ 'p:package',
        \ 'i:imports:1',
        \ 'c:constants',
        \ 'v:variables',
        \ 't:types',
        \ 'n:interfaces',
        \ 'w:fields',
        \ 'e:embedded',
        \ 'm:methods',
        \ 'r:constructor',
        \ 'f:functions'
    \ ],
    \ 'sro' : '.',
    \ 'kind2scope' : {
        \ 't' : 'ctype',
        \ 'n' : 'ntype'
    \ },
    \ 'scope2kind' : {
        \ 'ctype' : 't',
        \ 'ntype' : 'n'
    \ },
    \ 'ctagsbin'  : 'gotags',
    \ 'ctagsargs' : '-sort -silent'
\ }


"=============================================== YouCompleteMe =================================================
let g:ycm_min_num_of_chars_for_completion = 2      " 触发标识符补全的最小字符数
let g:ycm_max_num_identifier_candidates = 10       " 标识符补全的最大候选项数量，0 表示没有限制
let g:ycm_max_num_candidates = 10                  " 设置语义补全的最大候选项数量，0 表示没有限制
let g:ycm_auto_trigger=1                           " 自动打开补全
" 文件类型黑名单，打开这些类型文件时会关闭 YCM
let g:ycm_filetype_blacklist = {'tagbar' : 1, 'qf' : 1, 'notes' : 1, 'markdown' : 1, 'unite' : 1, 'text' : 1, 'vimwiki' : 1, 'pandoc' : 1, 'infolog' : 1, 'mail' : 1}
" 语义补全黑名单，打开这些类型文件时会关闭 YCM，但标识符补全仍可用
let g:ycm_filetype_specific_completion_to_disable={'gitcommit': 1}
" 对特定文件类型禁用文件路径补全
let g:ycm_filepath_blacklist = {'html' : 1, 'jsx' : 1,'xml' : 1}
let g:ycm_show_diagnostics_ui = 1                  " 开启显示诊断信息功能
let g:ycm_error_symbol = '>>'                      " 设置错误标志为 >>
let g:ycm_warning_symbol = '>*'                    " 设置警告标志为 *>
let g:ycm_enable_diagnostic_signs = 1              " 在代码中高亮显示诊断内容
let g:ycm_enable_diagnostic_highlighting = 1       " 高亮显示代码中与诊断信息有关的文本或代码
let g:ycm_echo_current_diagnostic = 1              " 当光标移到所在行时显示诊断信息
let g:ycm_open_loclist_on_ycm_diags = 1            " 在强制编译后自动打位置列表并用诊断信息填充，所谓位置列表就是标出各错误或警告对应在哪些行的小窗口，可以实现直接跳转到错误行
let g:ycm_complete_in_comments = 1                 " 补全功能在注释中同样有效
let g:ycm_complete_in_strings = 1                  " 打开字符串自动补全功能
let g:ycm_collect_identifiers_from_tags_files = 1  " 开启 tags 补全引擎，只支持 ctags
let g:ycm_server_python_interpreter = ''           " 指定特定的 Python 解释器，默认为空表示在系统上搜索适当的 Python 解释器
let g:ycm_global_ycm_extra_conf = ''               " 设置 .ycm_extra_conf.py 的全局路径，避免每次都需要复制到当前目录, 若为空则每次都需复制 .ycm_extra_conf.py 文件到当前目录
" 设置语义触发器的关键字
let g:ycm_semantic_triggers = {'c' : ['->', '.'], 'cs,java,javascript,typescript,d,python,perl6,scala,vb,elixir,go' : ['.'],'ruby' : ['.', '::']}
let g:ycm_use_ultisnips_completer = 1              " 启用 ultisnips 补全
let g:ycm_autoclose_preview_window_after_completion = 1    " 选中补全选项后自动关闭预览窗口
let g:ycm_confirm_extra_conf = 1                   " 允许自动加载 .ycm_extra_conf.py
let g:ycm_seed_identifiers_with_syntax=1           " 使用 VIM 的语法标识符来建立标识符数据库


"=============================================== vim-easymotion 快速移动光标
let g:EasyMotion_smartcase = 1
map <Leader><Leader>j <Plug>(easymotion-j)
map <Leader><Leader>k <Plug>(easymotion-k)


"=============================================== Syntastic
let g:syntastic_error_symbol = '>>'
let g:syntastic_warning_symbol = '>*'
let g:syntastic_check_on_open = 0                     " 打开文件时不会自动突出显示错误
let g:syntastic_check_on_wq = 0                       " 保存退出时不会自动突出显示错误
let g:syntastic_enable_highlighting = 1               " 高亮错误
let g:syntastic_python_checkers = ['pyflakes']        " 使用 pyflakes 检查 Python 语法，速度比 pylint 快
let g:syntastic_go_checkers = ['golint']              " 使用 golint 检查 Golang 语法
" 修改高亮的背景色, 适应主题
highlight SyntasticErrorSign guifg=white guibg=black

" to see error location list
let g:syntastic_always_populate_loc_list = 1          " 总是打开 Location List
let g:syntastic_auto_loc_list = 1                     " 自动打开/自动关闭 Locaton List
let g:syntastic_loc_list_height = 5                   " Locaton List 窗口高度
function! ToggleErrors()
    let old_last_winnr = winnr('$')
    lclose
    if old_last_winnr == winnr('$')
        " Nothing was closed, open syntastic error location panel
        Errors
    endif
endfunction
nnoremap <Leader>s :call ToggleErrors()<cr>
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*


"=============================================== Powerline
let g:airline_powerline_fonts = 1
let g:minBufExplForceSyntaxEnable = 1
let g:Powerline_symbols= "fancy"
set t_Co=256
set guifont=Source\ Code\ Pro\ for\ Powerline:h11


"=============================================== vim-go
let g:go_fmt_command = "goimports"                    " 格式化将默认的 gofmt 替换
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"
let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1
let g:godef_split=2
