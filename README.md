### Hexlet tests and linter status:

[![Actions Status](https://github.com/jobsboris27/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/jobsboris27/go-project-242/actions)

# hexlet-path-size

Утилита для подсчёта размера файла или директории. Поддерживает рекурсивный подсчёт, человеко-читаемый вывод и учёт скрытых файлов.

## Демонстрация работы

![Local GIF](./demo.gif)

## Установка

```bash
git clone https://github.com/jobsboris27/go-project-242.git
cd hexlet-path-size
make build
```

## Использование

hexlet-path-size [флаги] <путь>

### Примеры

Размер файла:

```bash
hexlet-path-size ./file.txt
```

Размер директории:

```bash
hexlet-path-size ./my-folder
```

Рекурсивный подсчёт (суммарно для всех вложенных директорий):

```bash
hexlet-path-size -r ./my-folder
```

Вывод в удобном формате (KB, MB, GB):

```bash
hexlet-path-size -H ./file.iso
```

Учёт скрытых файлов и директорий:

```bash
hexlet-path-size -a ./my-folder
```

Комбинация флагов:

```bash
hexlet-path-size -r -H -a ./my-folder
```

## Флаги

| Флаг          | Алиас | Описание                                              |
| ------------- | ----- | ----------------------------------------------------- |
| `--recursive` | `-r`  | Считать размер директорий рекурсивно                  |
| `--human`     | `-H`  | Вывод в человеко-читаемом формате (KB, MB, GB и т.д.) |
| `--all`       | `-a`  | Учитывать скрытые файлы и директории                  |

## Пример вывода

```bash
$ hexlet-path-size -r -H ./my-folder
12.3MB   ./my-folder
```
