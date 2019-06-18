---
title:  Go 开发工具 vim配置
date: 2018-12-06T00:00:00+08:00
---

## 粘贴代码块中的配置  $HOME/.vimrc 文件

```

" ====================== vim 配色设置 =====================
  set encoding=utf-8
  syntax enable
  set background=dark
  colorscheme solarized
" 设置字体
  set guifont=Monaco:h14
" solarized主题设置在终端下的设置
  let g:solarized_termcolors=256

" ===================== vim 原生配置  =====================
" 设定默认解码
  set fenc=utf-8
" history文件中需要记录的行数 
  set history=100
" 在处理未保存或只读文件的时候，弹出确认 
  set confirm
"在编辑过程中，在右下角显示光标位置的状态行
  set ruler
" 隐藏滚动条
  set guioptions-=r
  set guioptions-=L
  set guioptions-=b
" 开启语法高亮
  syntax on
" 设置不折行
  set nowrap
" 设置以unix的格式保存文件
  set fileformat=unix
" 设置C样式的缩进格式
  set cindent
" 距离顶部和底部5行
  set scrolloff=5
" 命令行行高
  set cmdheight=2
" 显示状态栏
  set laststatus=2
" 启动的时候不显示那个援助乌干达儿童的提示 
  set shortmess=atI 
" 通过使用: commands命令，告诉我们文件的哪一行被改变过
  set report=0 
" 不让vim发出讨厌的滴滴声
  set noerrorbells 

" 高亮搜索项
  set hlsearch
" 保存全局变量 
  set viminfo+=!
" 带有如下符号的单词不要被换行分割 
  set iskeyword+=_,$,@,%,#,-
" 不允许扩展table
  set noexpandtab
" 文件在Vim之外修改过，自动重新读入
  set autoread
" 突出显示当前行
  set cursorline
" 突出显示当前列
  set cursorcolumn
" 定义快捷键的前缀，即<Leader>
  let mapleader=";"
" 让配置变更立即生效
  autocmd BufWritePost $MYVIMRC source $MYVIMRC

" 开启实时搜索功能
  set incsearch
" 搜索时大小写不敏感
  set ignorecase
" vim 末行模式，命令自动补全，在状态栏有显示。按 TAB 键
  set wildmenu
" 设置tab键为4个空格
  set tabstop=4
  set shiftwidth=4
" 敲入tab键时实际占有的列数 并且将tab转换为空格
  set softtabstop=4 expandtab
" 设置匹配模式，类似当输入一个左括号时会匹配相应的右括号
  set showmatch
" 显示行号
  set number
" 使退格键有效
  set backspace=2
  set backspace=indent,eol,start
" 解决粘贴自动缩进
  set pastetoggle=<F2>
" 高亮字符，让其不受100列限制
  :highlight OverLength ctermbg=red ctermfg=white guibg=red guifg=white
  :match OverLength '\%101v.*'

" 状态行颜色
  highlight StatusLine guifg=SlateBlue guibg=Yellow
  highlight StatusLineNC guifg=Gray guibg=White 
" 增强模式中的命令行自动完成操作 
  set wildmenu

" Tmux 下vim 正确显示
  if exists('$TMUX')
      set term=screen-256color
  endif

"===================== vim 文件配置  =====================
" 不要备份文件（根据自己需要取舍)
  set nobackup
" 不要生成swap文件，当buffer被丢弃的时候隐藏它
  setlocal noswapfile
  set bufhidden=hide
" 字符间插入的像素行数目
  set linespace=0
" 文件编码
  set fenc=utf-8
" 缩进
  let g:indent_guides_auto_colors = 1
  let g:indent_guides_start_level = 1
  let g:indent_guides_guide_size = 1
  let g:indent_guides_enable_on_vim_startup = 1

"===================== vim 插件配置  =====================

  set nocompatible
  filetype off

" 设置vundle 初始化路径
  set rtp+=~/.vim/bundle/Vundle.vim
  call vundle#begin()

" ale 与 Syntastic 冲突
  let g:ale_emit_conflict_warnings = 0

" 设置Vundle管理vim插件 这是必须的
  Plugin 'VundleVim/Vundle.vim'

" 配色插件 有点像sublimetext
  Plugin 'tomasr/molokai'

" 显示末尾空格的插件
  Plugin 'ShowTrailingWhitespace'
" 安装solarized
  Plugin 'altercation/vim-colors-solarized'

" 安装Your Complete Me
  Plugin 'Valloric/YouCompleteMe'

" YouCompleteMe 配置
  let g:ycm_register_as_syntastic_checker = 0
  let g:ycm_min_num_of_chars_for_completion = 10
  let g:ycm_min_num_identifier_candidate_chars = 10
  let g:ycm_filetype_whitelist = { 'cpp': 1 }
  let g:ycm_filetype_specific_completion_to_disable = { 'cpp': 1 }
  let g:ycm_cache_omnifunc = 0

" 安装NerdTree 插件
  Plugin 'scrooloose/nerdtree'

" NerdTree 配置
  map <F8> :NERDTreeToggle<CR>

  let NERDTreeIgnore=['\.o$','\.a$', '\.pyc$', '\.taghl$','\~$', 'cscope\.', 'tags$', '\.bak$', '\.php\~$']
  let NERDTreeChDirMode = 2
  let NERDTreeWinSize = 20
  let NERDTreeShowBookmarks = 1
  autocmd VimEnter * NERDTree


" 安装Emmet插件
  Plugin 'mattn/emmet-vim'

" 安装clipboard插件 方便复制粘贴
  

" 安装Tagbar插件 生成大纲啊，选中快速跳转到目标位置 系统必须安装Exuberant ctags
" sudo apt install exuberant-ctags && sudo yum install exuberant-ctags
  Plugin 'majutsushi/tagbar'

" Tagbar 配置
  map <F6> :TagbarToggle<CR>
  let g:tagbar_sort = 0
  let g:tagbar_width = 20

" 添加PHP定义到tagbar
  let g:tagbar_type_php = {
    \ 'kinds' : [
        \ 'i:interfaces:1',
        \ 'c:classes:1',
        \ 'd:constant definitions:1:0',
        \ 'f:functions',
        \ 'v:variables:1:0',
        \ 'j:javascript functions:1',
    \ ],
\ }

" 完美的vim缩进提示线  :IndentLinesToggle  命令切换线条打开和关闭
  Plugin 'Yggdroot/indentLine'
" indentLine 配置 定制隐藏颜色
  let g:indentLine_color_term = 239
" 背景色
  let g:indentLine_bgcolor_term = 202
  let g:indentLine_bgcolor_gui = '#FF5F00'
" 更改缩进字符 这些字符只适用于UTF-8编码的文件
  let g:indentLine_char = '┆'


" gitv插件 Vim来查看Git的详细提交信息
  Plugin 'gregsexton/gitv'
" vim-gitgutter 管理项目
  Plugin 'airblade/vim-gitgutter'

" 强大的fzf 快速搜索
  Plugin 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
  Plugin 'junegunn/fzf.vim'

" fzf 配置 
  command! -bang -nargs=* Rg
  \ call fzf#vim#grep(
  \   'rg --column --line-number --no-heading --color=always '.shellescape(<q-args>), 1,
  \   <bang>0 ? fzf#vim#with_preview('up:60%')
  \           : fzf#vim#with_preview('right:50%:hidden', '?'),
  \   <bang>0)

  map <leader>ff :<C-u>Files<CR>
  map <leader>b :<C-u>Buffers<CR>
  map <leader>t :<C-u>Tags <C-R><C-W><CR>
  map <C-T><C-T> :<C-u>Tags

  map <C-H><C-H> eb :Ag <C-R><C-W><CR>

" Ack插件
  Plugin 'mileszs/ack.vim'
" Ack配置
  let g:ackprg = 'ag --nogroup --nocolor --column'

" 错误检查
  Plugin 'w0rp/ale'

" 错误检查插件
  Plugin 'vim-syntastic/syntastic'
" syntastic 配置
  set statusline+=%#warningmsg#
  set statusline+=%{SyntasticStatuslineFlag()}
  set statusline+=%*

  let g:syntastic_always_populate_loc_list = 1
  let g:syntastic_auto_loc_list = 1

  let g:syntastic_check_on_open = 0
  let g:syntastic_check_on_wq = 0
  let g:syntastic_php_checkers = ['php']
  let g:syntastic_quiet_messages = { "type": "style" }
  let g:syntastic_enable_signs = 1
  let g:syntastic_ignore_files = ['\m\c\.cc$', '\m\c\.h$']

" 注释插件
  Plugin 'tpope/vim-commentary'

" 对文本中的多个字串以不同的颜色高亮显示
  Plugin 'vim-scripts/Mark--Karkat'

" snipmate 代码片段 这个暂时挖坑 这货还老与YouCompleteMe冲突

" statusline 状态行
  Plugin 'vim-airline/vim-airline'

" airline 配置
  let g:airline#extensions#tabline#enabled = 1
  let g:airline#extensions#tabline#left_sep = ' '
  let g:airline#extensions#tabline#left_alt_sep = '|'

" 语言脚本配置
  Plugin 'plasticboy/vim-markdown'

" PHP complete 自动补全
  Plugin 'shawncplus/phpcomplete.vim'
 
  Plugin 'rayburgemeestre/phpfolding.vim'

" go 配置
" 安装vim-go
  Plugin 'fatih/vim-go'
  
" go bin
  let g:go_bin_path = "$GOBIN"
  let g:go_fmt_command = "goimports"
  let g:go_metalinter_autosave = 1
  let g:go_metalinter_autosave_enabled = ['errcheck']
  let g:go_metalinter_deadline = "30s"
  let g:go_list_height = 20
  
 
  let g:go_highlight_extra_types = 1
  let g:go_highlight_functions = 1
  let g:go_highlight_methods = 1
  let g:go_highlight_fields = 1
  let g:go_highlight_interfaces = 0
  let g:go_highlight_structs = 0
  let g:go_highlight_operators = 0
  let g:go_highlight_build_constraints = 1
  let g:go_highlight_format_strings = 1
  let g:go_auto_type_info = 0

  let g:go_guru_scope = ["maid", "Gout"]

" 其他插件
  Plugin 'tpope/vim-repeat'

  Plugin 'Lokaltog/vim-easymotion'

" extend %
" 默认的% 只能匹配简单的比如括号, 这个扩展了一些
  Plugin 'matchit.zip'
  Plugin 'tpope/vim-surround'

" 显示命令行颜色
  Plugin 'powerman/vim-plugin-AnsiEsc'

" -------------------MBF-------------------------
  let g:miniBufExplMapWindowNavVim = 1
  let g:miniBufExplMapWindowNavArrows = 1
  let g:miniBufExplMapCTabSwitchBufs = 1
  let g:miniBufExplModSelTarget = 1
  let g:miniBufExplMaxHeight = 1
" MiniBufExpl Colors
  hi MBEVisibleActive guifg=#A6DB29 guibg=fg
  hi MBEVisibleChangedActive guifg=#F1266F guibg=fg
  hi MBEVisibleChanged guifg=#F1266F guibg=fg
  hi MBEVisibleNormal guifg=#5DC2D6 guibg=fg
  hi MBEChanged guifg=#CD5907 guibg=fg
  hi MBENormal guifg=#808080 guibg=fg
  let g:miniBufExplorerMoreThanOne=0
  let g:miniBufExplSplitBelow=0 " Put new window above

" ----------------------taglist-----------------------
  let Tlist_Auto_Open=1
  let Tlist_Show_One_File = 1
  let Tlist_Exit_OnlyWindow = 1
  let Tlist_Use_Right_Window = 1


  call vundle#end()            " required
  filetype plugin indent on    " required

  ```
