module LogLevels

let message (logLine: string): string = 
    let parts = logLine.Split ':'
    let msg = parts[1]
    msg.Trim()

let logLevel(logLine: string): string = 
    let level = logLine.[1+logLine.IndexOf '[' .. logLine.IndexOf ']'-1]
    level.ToLower()

let reformat(logLine: string): string = 
    message(logLine) + " (" + logLevel(logLine) + ")"

