<template>
  <div id="terminal" class="hello">
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Terminal } from 'xterm';
import 'xterm/dist/xterm.css';

@Component
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;
  protected mounted(): void {
      const t = new Terminal();
      const serverURL = location.host || 'localhost';
      const conn = new WebSocket(`ws://${serverURL}/ws`);
      const el = document.getElementById('terminal');
      if (el == null) {
          throw Error('element didn\'t exist');
      }

      t.open(el);
      t.addDisposableListener('key', (key, ev) => {
          const printable = !ev.altKey && !ev.altGraphKey && !ev.ctrlKey && !ev.metaKey;

          if (ev.code === 13) {
              t.write('\r\n$ ');
          } else if (ev.code === 8) {
              // Do not delete the prompt
              if (t.x > 2) {
                  t.write('\b \b');
              }
          } else if (printable) {
              // t.write(key);
              conn.send(key);
          }
      })

      conn.onclose = (event:any): void => {
          t.writeln("")
          t.writeln("[closed]")
      }
      conn.onmessage = (event:any): void => {
          t.write(event.data)
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
