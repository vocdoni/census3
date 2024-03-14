import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.prod.js'

// The API URL can be changed with the setAPI function in the browser console 
// with the admin password and the new URL as parameters
let apiURL = "https://census3-dev.vocdoni.net/api";
// admin password is a hashed password with SHA-256 that allows to change the 
// API URL using the setAPI function in the browser console, by default it is
// the hash of the string "admin"
// const adminPassword = "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918";
const adminPassword = "9d92eb71f8603415139843fa11b7d823ae8062481a41b051d520404f2c90e4c7";

async function sha256(inputString) {
    const utf8 = new TextEncoder().encode(inputString);
    const hashBuffer = await crypto.subtle.digest('SHA-256', utf8);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map((bytes) => bytes.toString(16).padStart(2, '0')).join('');
}

window.setAPI = async function(password, url) {
    const hashPassword = await sha256(password);
    if (hashPassword !== adminPassword) {
        throw new Error("Invalid password");
    }
    apiURL = url;
    console.log(`API URL set to ${apiURL}`);
}

console.log(`Using API URL: ${apiURL}. It can be changed with the setAPI function in the browser console with the admin password and the new URL as parameters`);

const app = createApp({
    data() {
        return {
            selectedToken: null,
            tokens: [],
            loading: false,
            fileURI: null
        }
    },
    async created() {
        await this.loadTokens();
    },
    template: `
    <div id="app" class="container">
        <h1>👾 hold3rs</h1>
        <p>A simple web app to get the holders of a token in CSV from <a href="https://github.com/vocdoni/census3" target="_blank">Census3</a>.</p>
        <div class="selector-container">
            <label for="tokens"><b>Select a Token</b></label>
            <select id="tokens" v-model="selectedToken" @change="reset">
                <option v-for="(token, index) in tokens" :key="index" :value="token">
                    {{ token.name }} ({{ token.symbol }})
                </option>
            </select>
        </div>
        <button type="button" @click="launchCSV" :disabled="!selectedToken || loading" class="generate-btn">
            Generate CSV 📝
        </button>
        <span v-if="loading" class="loading-text">
            <div class="spinner"></div>
            Loading...
        </span>
        <a v-if="fileURI" :href="fileURI" :download="selectedToken.symbol+'.csv'" class="download-link">Download CSV 📥</a>
    </div>
    `,
    methods: {
        async loadTokens() {
            try {
                const getTokensURI = `${apiURL}/tokens?pageSize=-1`;
                const response = await fetch(getTokensURI);
                const data = await response.json();
                this.tokens = data.tokens;
            } catch (e) {
                console.error("error getting the tokens", e);
                alert("Error getting the tokens :( Please try again later!");
            }
        },
        async launchCSV() {
            this.fileURI = null;
            let launchURI = `${apiURL}/tokens/${this.selectedToken.ID}/csv?chainID=${this.selectedToken.chainID}`;
            if (this.selectedToken.externalID) {
                launchURI += `&externalID=${this.selectedToken.externalID}`;
            }
            try {
                const response = await fetch(launchURI);
                const data = await response.json();
                this.loading = true;
                const csv = await this.getCSV(this.selectedToken.ID, data.queueID);
                this.loading = false;
                const blob = new Blob([csv], { type: 'text/csv' });
                this.fileURI = URL.createObjectURL(blob);
            } catch (e) {
                this.loading = false;
                console.error("error creating the csv", e);
                alert("Error creating the CSV :( Please try again later!");
            }
        },
        async getCSV(token, queueID) {
            const queueURI = `${apiURL}/tokens/${token}/csv/queue/${queueID}`;

            while (true) {
                const response = await fetch(queueURI);
                if (response.status === 204) {
                    await new Promise(resolve => setTimeout(resolve, 1000));
                    continue;
                }
                if (response.status === 200) return await response.text();
                const body = await response.text();
                throw new Error(`CSV failed: ${body}`);
            }
        },
        reset() {
            this.loading = false;
            this.fileURI = null;
        }
    }
});

app.mount('#app');