<script>
    import { onMount, onDestroy } from 'svelte';

    let mensajes = [];
    let texto = '';
    let autor = '';
    let intervalo;

    onMount(async () => {
        autor = localStorage.getItem('autor');
        if (!autor) {
            autor = prompt('¿Cuál es tu nombre?') || 'Anónimo';
            localStorage.setItem('autor', autor);
        }
        await cargarMensajes();
        intervalo = setInterval(cargarMensajes, 2000);
    });

    onDestroy(() => {
        clearInterval(intervalo);
    });

    async function cargarMensajes() {
        const res = await fetch('http://localhost:8080/messages');
        mensajes = await res.json();
    }

    async function enviarMensaje() {
        if (!texto.trim()) return;
        await fetch('http://localhost:8080/messages', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ author: autor, text: texto })
        });
        texto = '';
        await cargarMensajes();
    }

    async function eliminarMensaje(id) {
        await fetch(`http://localhost:8080/messages/${id}`, {
            method: 'DELETE'
        });
        await cargarMensajes();
    }

    function formatearHora(timestamp) {
        return new Date(timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
    }
</script>

<div class="chat-container">
    <div class="chat-header">
        <h2>Chat basico
        </h2>
    </div>

    <div class="mensajes">
        {#each mensajes as msg}
            <div class="burbuja {msg.Author === autor ? 'propio' : 'ajeno'}">
                {#if msg.Author !== autor}
                    <span class="autor">{msg.Author}</span>
                {/if}
                <p>{msg.Text}</p>
                <div class="meta">
                    <span class="hora">{formatearHora(msg.Timestamp)}</span>
                    {#if msg.Author === autor}
                        <button class="eliminar" on:click={() => eliminarMensaje(msg.ID)}>🗑</button>
                    {/if}
                </div>
            </div>
        {/each}
    </div>

    <div class="input-bar">
        <input
            bind:value={texto}
            placeholder="Escribí un mensaje..."
            on:keydown={(e) => e.key === 'Enter' && enviarMensaje()}
        />
        <button on:click={enviarMensaje}>➤</button>
    </div>
</div>

<style>
    * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    .chat-container {
        display: flex;
        flex-direction: column;
        height: 100vh;
        max-width: 500px;
        margin: 0 auto;
        font-family: sans-serif;
        background: #e5ddd5;
    }

    .chat-header {
        background: #075e54;
        color: white;
        padding: 16px;
        text-align: center;
    }

    .mensajes {
        flex: 1;
        overflow-y: auto;
        padding: 16px;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .burbuja {
        max-width: 75%;
        padding: 8px 12px;
        border-radius: 8px;
        position: relative;
    }

    .propio {
        background: #dcf8c6;
        align-self: flex-end;
        border-bottom-right-radius: 0;
    }

    .ajeno {
        background: white;
        align-self: flex-start;
        border-bottom-left-radius: 0;
    }

    .autor {
        font-size: 0.75em;
        font-weight: bold;
        color: #075e54;
        display: block;
        margin-bottom: 4px;
    }

    p {
        font-size: 0.95em;
    }

    .meta {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        gap: 6px;
        margin-top: 4px;
    }

    .hora {
        font-size: 0.7em;
        color: #888;
    }

    .eliminar {
        background: none;
        border: none;
        cursor: pointer;
        font-size: 0.8em;
        padding: 0;
    }

    .input-bar {
        display: flex;
        gap: 8px;
        padding: 12px;
        background: #f0f0f0;
    }

    input {
        flex: 1;
        padding: 10px 14px;
        border-radius: 20px;
        border: none;
        outline: none;
        font-size: 0.95em;
    }

    button {
        background: #075e54;
        color: white;
        border: none;
        border-radius: 50%;
        width: 42px;
        height: 42px;
        font-size: 1.1em;
        cursor: pointer;
    }
</style>