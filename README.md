# task-cli

https://roadmap.sh/projects/task-tracker

---

## Run

```sh
go run main.go [commands]
```

## Build

```sh
go build
```

## Example Commands

### Adding a new task

```sh
./task-cli add "Buy groceries"
```

### Updating and deleting tasks

```sh
./task-cli update 1 "Buy groceries and cook dinner"
```

```sh
./task-cli delete 1
```
### Marking a task as in progress or done

```sh
./task-cli mark-in-progress 1
```

```sh
./task-cli mark-done 1
```
### Listing all tasks

```sh
./task-cli list
```

### Listing tasks by status

```sh
./task-cli list done
```

```sh
./task-cli list todo
```

```sh
./task-cli list in-progress
```
