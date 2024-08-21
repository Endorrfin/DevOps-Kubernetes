import fs from 'fs'

fs.appendFile('my-file.txt', 'File created Node.js', (err) => {
    if (err) throw err
    console.log('✅ File seved!')
})

setTimeout(() => console.info("🟥 The End!"), 10000)