# Go Logger Init Log With Slogger
una implementacion simple para el paquete de slog de golang

## Installation

```bash
# Go Modules
require github.com/guillospy92/logger
```

## Usage
### import package
```
import "github.com/guillospy92/clientHttp/logger"
```


### calls
logger.Log() return  *slog.Logger

## simple example
logger.Log().With(...).Info("message ...")


## examples add context
```
ctx = logger.AppendCtx(ctx, slog.String("uuid", "aefd-fddv-edft"))
logger.Log().LogAttrs(ctx, slog.LevelError, logMessage, attr...)
```

## Utils
```
declare the environment variable LOGGER_SAVE_FILE if you want to save logs to files
compatible and saves the trace of the library github.com/pkg/errors
```



