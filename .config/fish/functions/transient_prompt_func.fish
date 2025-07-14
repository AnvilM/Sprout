function transient_prompt_func
    set_color green
    if test $transient_status -ne 0
        set_color red
    end
    echo -n "❯ "
    set_color normal
end
