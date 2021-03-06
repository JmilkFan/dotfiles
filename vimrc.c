"=============================================== Vunble =================================================
" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" alternatively, pass a path where Vundle should install plugins
"call vundle#begin('~/some/path/here')

" let Vundle manage Vundle, required
Plugin 'VundleVim/Vundle.vim'

" The following are examples of different formats supported.
" Keep Plugin commands between vundle#begin/end.

Plugin 'altercation/vim-colors-solarized'
Plugin 'altercation/solarized'
Plugin 'scrooloose/nerdtree'
Plugin 'jistr/vim-nerdtree-tabs'
Plugin 'kien/ctrlp.vim'
Plugin 'majutsushi/tagbar'
Plugin 'Valloric/YouCompleteMe'
Plugin 'scrooloose/syntastic'
Plugin 'vim-scripts/TaskList.vim'
Plugin 'easymotion/vim-easymotion'
Plugin 'scrooloose/nerdcommenter'
Plugin 'Lokaltog/powerline', {'rtp': 'powerline/bindings/vim/'}
Plugin 'powerline/fonts'

" Plugin 'SirVer/ultisnips'
" Plugin 'tpope/vim-surround'
" Plugin 'plasticboy/vim-markdown'
" Plugin 'vim-scripts/indentpython.vim'
" Plugin 'tpope/vim-fugitive'
" Plugin 'christoomey/vim-run-interactive'
" Plugin 'tmhedberg/SimpylFold'

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required
" To ignore plugin indent changes, instead use:
filetype plugin on
"
" Brief help
" :PluginList       - lists configured plugins
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate
" :PluginSearch foo - searches for foo; append `!` to refresh local cache
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal
"
" see :h vundle for more details or wiki for FAQ
" Put your non-Plugin stuff after this line


"=============================================== General Config =================================================
set nocompatible                          " be iMproved, required
filetype off                              " required
set number                                " Line numbers are good
set ruler                                 " Show the cursor position all the time
set backspace=indent,eol,start            " Allow backspace in insert mode
set fileencodings=utf-8,gbk               " Set encoding of files
set history=1000                          " Number of things to remember in history
set showcmd                               " Show incomplete cmds down the bottom
set showmode                              " Show current mode down the bottom
set showmatch                             " Show matching brackets
set gcr=a:blinkon0                        " Disable cursor blink
set novisualbell                          " No sounds
set noerrorbells                          " No noise
set autoread                              " Reload files changed outside vim
set statusline+=%{fugitive#statusline()}  " Git Hotness
set list listchars=tab:>.,trail:.         " Display tabs and trailing spaces visually
set linebreak                             " Wrap lines at convenient points
set wildmenu                              " Command auto complete
set shortmess=atI
set statusline=\ %<%F[%1*%M%*%n%R%H]%=\ %y\ %0(%{&fileformat}\ %{&encoding}\ %c:%l/%L%)\
set nobackup
set confirm
set nowb
set tabstop=4
set shiftwidth=4
set textwidth=79                          " Make it obvious where 79 characters is
highlight ColorColumn ctermbg=gray
set colorcolumn=79
set numberwidth=4
set fileformat=unix
set expandtab
set list
set cmdheight=1
set ignorecase smartcase
au WinLeave * set nocursorline nocursorcolumn   " Highlight current line
au WinEnter * set cursorline cursorcolumn
set cursorline cursorcolumn
set clipboard=unnamed

"split navigations
set splitbelow
set splitright
nnoremap <C-J> <C-W><C-J>
nnoremap <C-K> <C-W><C-K>
nnoremap <C-L> <C-W><C-L>
nnoremap <C-H> <C-W><C-H>

" Persistent Undo
set undodir=~/.vim/backups
set undofile

" Search Options
set incsearch                   " Find the next match as we type the search
set hlsearch                    " Hilight searches by default
set viminfo='100,f1             " Save up to 100 marks, enable capital marks

" Indentation
set autoindent
set smartindent
set smarttab
set shiftwidth=4
set softtabstop=4
set tabstop=4
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


"=============================================== NERD Tree =================================================
let NERDChristmasTree=0
let NERDTreeWinSize=35
let NERDTreeChDirMode=2
let NERDTreeIgnore=['\~$', '\.pyc$', '\.swp$']
let NERDTreeShowBookmarks=1
let NERDTreeWinPos="left"
" Automatically open a NERDTree if no files where specified
autocmd vimenter * if !argc() | NERDTree | endif
" Close vim if the only window left open is a NERDTree
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTreeType") && b:NERDTreeType == "primary") | q | endif
" Open a NERDTree
nmap <leader>n :NERDTreeToggle<CR>


"=============================================== Ctrlp =================================================
set wildignore+=*/tmp/*,*.so,*.swp,*.zip,*.png,*.jpg,*.jpeg,*.gif  " Ignore for MacOSX/Linux
let g:ctrlp_custom_ignore = {
    \ 'dir':  '\v[\/]\.(git|hg|svn|rvm)$',
    \ 'file': '\v\.(exe|so|dll|zip|tar|tar.gz|pyc)$',
    \ }
let g:ctrlp_match_window = 'bottom,order:btt,min:1,max:10,results:20'
let g:ctrlp_max_height = 30
"let g:ctrlp_user_command = [
"    \ '.git', 'cd %s && git ls-files . -co --exclude-standard',
"    \ 'find %s -type f'
"    \ ]
if executable('ag')
  " Use Ag over Grep
  set grepprg=ag\ --nogroup\ --nocolor
  " Use ag in CtrlP for listing files.
  let g:ctrlp_user_command = 'ag %s -l --nocolor -g ""'
  " Ag is fast enough that CtrlP doesn't need to cache
  let g:ctrlp_use_caching = 0
endif
let g:ctrlp_working_path_mode=0
let g:ctrlp_match_window_bottom=1
let g:ctrlp_max_height=15
let g:ctrlp_match_window_reversed=0
let g:ctrlp_mruf_max=500
let g:ctrlp_follow_symlinks=1
let g:ctrlp_map = '<leader>p'
let g:ctrlp_cmd = 'CtrlP'
nmap <leader>f :CtrlPMRU<CR>


"=============================================== YouCompleteMe =================================================
let g:ycm_global_ycm_extra_conf='~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp/ycm/.ycm_extra_conf.py'
let g:ycm_autoclose_preview_window_after_completion=1
let g:syntastic_always_populate_loc_list = 1
let g:ycm_confirm_extra_conf = 0
let g:ycm_complete_in_comments = 1
let g:ycm_complete_in_strings = 1
let g:ycm_collect_identifiers_from_comments_and_strings = 0
let g:ycm_seed_identifiers_with_syntax=1
let g:ycm_cache_omnifunc=0
let g:ycm_min_num_of_chars_for_completion=2
let g:ycm_collect_identifiers_from_tags_files=1
let g:ycm_error_symbol = '>>'
let g:ycm_warning_symbol = '>*'
nmap <leader>gf :YcmCompleter GoToDefinition<CR>
nmap <leader>gl :YcmCompleter GoToDeclaration<CR>
nmap <leader>g :YcmCompleter GoToDefinitionElseDeclaration<CR>


"=============================================== Tagbar =================================================
let g:tagbar_width=35
let g:tagbar_autofocus=1
let g:tagbar_ctags_bin ='/usr/local/bin/ctags'
nmap <leader>c :TagbarToggle<CR>


"=============================================== Syntastic =================================================
" configure syntastic syntax checking to check on open as well as save
" let g:syntastic_check_on_open=1 " Using this option will be slow to open the file
let g:syntastic_html_tidy_ignore_errors=[" proprietary attribute \"ng-"]
let g:syntastic_always_populate_loc_list = 1
let g:syntastic_auto_loc_list = 1
let g:syntastic_check_on_wq = 0
set statusline+=%#warningmsg#
set statusline+=%{SyntasticStatuslineFlag()}
set statusline+=%*


"=============================================== Solarized =================================================
syntax enable
set background=dark
let g:solarized_termcolors=16
let g:solarized_visibility='high'
let g:solarized_contrast='high'
try
  colorscheme solarized
catch /^Vim\%((\a\+)\)\=:E185/
endtry


"=============================================== Powerline =================================================
let g:airline_powerline_fonts = 1
let g:minBufExplForceSyntaxEnable = 1
let g:Powerline_symbols= "fancy"
set laststatus=2                          " Always show status line
set encoding=utf-8
set t_Co=256
set guifont=Source\ Code\ Pro\ for\ Powerline:h11


"=============================================== flake8-vim =================================================
let g:PyFlakeDisabledMessages = 'E501'
let g:PyFlakeOnWrite = 1
let g:PyFlakeCheckers = 'pep8'
let g:PyFlakeAggressive = 0
let g:PyFlakeCWindow = 6
let g:PyFlakeSigns = 1
let g:PyFlakeSignStart = 1
nmap <leader>pt :PyFlakeToggle<CR>
nmap <leader>pf :PyFlake<CR>
nmap <leader>pa :PyFlakeAuto<CR>



"=============================================== vim-easymotion =================================================
let g:EasyMotion_smartcase = 1
map <Leader><leader>h <Plug>(easymotion-linebackward)
map <Leader><Leader>j <Plug>(easymotion-j)
map <Leader><Leader>k <Plug>(easymotion-k)
map <Leader><leader>l <Plug>(easymotion-lineforward)
map <Leader><leader>. <Plug>(easymotion-repeat)
