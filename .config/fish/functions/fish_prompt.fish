# Prompt
function fish_prompt
    set_color black
    echo -n "╭──"
    set_color normal
    echo -n ""

    prompt_open_bracket
    prompt_status_circle
    set_color normal
    prompt_close_bracket

    prompt_open_bracket
    set_color green
    prompt_username
    set_color normal
    prompt_close_bracket

    prompt_open_bracket
    set_color blue
    prompt_current_directory
    set_color normal
    prompt_close_bracket

    prompt_git_status
    set_color normal

    echo -en "\n"

    set_color black
    echo -n "╰──"
    set_color normal
    echo -n ""

    echo -n "❯ "
end




# Right prompt
function fish_right_prompt
    tput sc; tput cuu1; tput cuf 2
    
    prompt_open_bracket
    set_color yellow
    prompt_command_execution_time
    set_color normal
    prompt_close_bracket

    tput rc
end




# Open bracket
function prompt_open_bracket
set_color black
echo -n "["
set_color normal
end




# Close bracket
function prompt_close_bracket
set_color black
echo -n "]"
set_color normal
end




# Status circle
function prompt_status_circle
    if test $transient_status -eq 0
        set_color green
        echo -n '●'
    else
        set_color red
        echo -n '●'
    end
    set_color normal
end




# Username
function prompt_username
    echo -n (whoami)
end




# Current directory
function prompt_current_directory
    set cwd (pwd)
    set home $HOME

    if string match -q "$home*" "$cwd"
        set rel_path "~"(string sub --start (math (string length -- $home) + 1) -- $cwd)
        echo -n $rel_path
    else
        echo -n $cwd
    end
end




# Command execution time 
function prompt_command_execution_time
    echo -n (math --scale 2 "$CMD_DURATION" / 1000)s
end




# Git info
function prompt_git_status
    if not git rev-parse --git-dir >/dev/null 2>&1
        return
    end

    prompt_open_bracket

    set -l branch (git symbolic-ref --short HEAD 2>/dev/null; or git rev-parse --short HEAD 2>/dev/null)
    if test -z "$branch"
        return
    end

    set -l git_status (git status --porcelain 2>/dev/null)
    set -l untracked 0
    set -l staged 0
    set -l modified 0

    for line in $git_status
        set -l index (string sub -l 1 $line)
        set -l worktree (string sub -s 2 -l 1 $line)

        if string match -qr '^\?\?' $line
            set untracked (math $untracked + 1)
        else if string match -qr '^[AMDR]' $line
            if string match -qr '^[AMDR][M]' $line
                set modified (math $modified + 1)
            else
                set staged (math $staged + 1)
            end
        else if string match -qr '^.[MD]' $line
            set modified (math $modified + 1)
        end
    end

    set -l branch_color green
    if test $untracked -gt 0 -o $staged -gt 0 -o $modified -gt 0
        set branch_color yellow
    end

    set_color $branch_color
    echo -n "$branch"
    set_color normal

    if test $untracked -gt 0
        set_color red
        echo -n " +$untracked"
        set_color normal
    end

    if test $staged -gt 0
        set_color green
        echo -n " ~$staged"
        set_color normal
    end

    if test $modified -gt 0
        set_color yellow
        echo -n " *$modified"
        set_color normal
    end
    prompt_close_bracket
end



