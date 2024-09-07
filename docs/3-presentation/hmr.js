// https://stackoverflow.com/a/69628635
export default function CustomHmr() {
  return {
    name: 'markdown-hmr',
    enforce: 'post',
    // HMR
    handleHotUpdate({ file, server }) {
      if (file.endsWith('.md')) {
        console.log('reloading markdown file...');

        server.ws.send({
          type: 'full-reload',
          path: '*'
        });
      }
    },
  }
}
