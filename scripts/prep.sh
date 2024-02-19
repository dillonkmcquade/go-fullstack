if ! command -v grep;
then
    echo "missing command: grep"
    exit 1
fi

if ! command -v rename;
then
    echo "missing command: rename"
    exit 1
fi

if ! command -v basename;
then
    echo "missing command: basename"
    exit 1
fi

if ! command -v dirname;
then
    echo "missing command: dirname"
    exit 1
fi

if ! command -v find;
then
    echo "missing command: find"
    exit 1
fi

if ! command -v xargs;
then
    echo "missing command: xargs"
    exit 1
fi

if ! command -v sed;
then
    echo "missing command: sed"
    exit 1
fi

dir_name=$(basename "$PWD")

echo "============ Prepping code repository ============"
echo ""
echo "Replacing in-text occurrences of 'go-fullstack' with $dir_name"
grep --exclude-dir=\*scripts\* -rl "go-fullstack" . | xargs sed -i "s/go-fullstack/$dir_name/g" && echo "Success."


echo "Replacing file and directory occurrences of 'go-fullstack' with $dir_name"
find . -type d -name "go-fullstack" -exec rename -ai "go-fullstack" "$dir_name" "{}" ";" && echo "Success."
