import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.prod.js'

let apiURL = "http://localhost:7788/api";
const adminPassword = "24dd8660eddc03d6da1b1168ef5f1ff8812bc024af9c7c80dd9bc81a2ca1ce90";

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
    template: `<div>
        <select id="tokens" v-model="selectedToken">
            <option v-for="(token, index) in tokens" :key="index" :value="token">{{ token.name }} ({{ token.symbol }})</option>
        </select>
        <button type="button" @click="launchCSV" :disabled="!selectedToken || loading">Generate CSV</button>
        <span v-if="loading">loading</span>
        <a v-if="fileURI" :href="fileURI" :download="selectedToken.symbol+'.csv'">Download CSV</a>
    </div>`,
    methods: {
        async loadTokens() {
            const getTokensURI = `${apiURL}/tokens?pageSize=-1`;
            const response = await fetch(getTokensURI);
            const data = await response.json();
            this.tokens = data.tokens;
        },
        async launchCSV() {
            this.fileURI = null;
            let launchURI = `${apiURL}/tokens/${this.selectedToken.ID}/csv?chainID=${this.selectedToken.chainID}`;
            if (this.selectedToken.externalID) {
                launchURI += `&externalID=${this.selectedToken.externalID}`;
            }
            const response = await fetch(launchURI);
            const data = await response.json();
            this.loading = true;
            const csv = await this.getCSV(this.selectedToken.ID, data.queueID);
            this.loading = false;
            const blob = new Blob([csv], { type: 'text/csv' });
            this.fileURI = URL.createObjectURL(blob);
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
        }
    }
});

app.mount('#app');