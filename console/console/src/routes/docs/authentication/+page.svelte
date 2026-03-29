<script lang="ts">
        import { baseUrl, copyText, copiedId } from '$lib/docs-store';
        import CodeBlock from '$lib/CodeBlock.svelte';
</script>

<svelte:head><title>Authentication — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Authentication</h1>
        <p class="text-sm text-[#666] mb-8 leading-relaxed">
                All API requests must include a valid API key. Keys are validated against a secure backend with a 60-second local cache for high performance.
        </p>

        <!-- Methods -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-4">Authentication methods</h2>
        <div class="space-y-4 mb-8">

                <!-- Method 1 -->
                <div class="border border-[#e5e5e5] rounded-2xl overflow-hidden">
                        <div class="flex items-center justify-between px-4 py-3 bg-[#f9f9f9] border-b border-[#e5e5e5]">
                                <div class="flex items-center gap-2">
                                        <span class="w-5 h-5 rounded-full bg-[#0f0f0f] text-white text-[10px] font-bold flex items-center justify-center">1</span>
                                        <span class="text-sm font-semibold text-[#333]">Authorization header</span>
                                        <span class="text-xs px-2 py-0.5 rounded-full bg-green-100 text-green-700 border border-green-200 font-medium">Recommended</span>
                                </div>
                        </div>
                        <div class="p-4 bg-[#f4f4f4]">
                                <p class="text-xs text-[#888] mb-3">Pass your API key in the standard HTTP Authorization header.</p>
                                <CodeBlock code="Authorization: Bearer sk-<your-api-key>" id="auth-bearer" />
                        </div>
                </div>

                <!-- Method 2 -->
                <div class="border border-[#e5e5e5] rounded-2xl overflow-hidden">
                        <div class="flex items-center gap-2 px-4 py-3 bg-[#f9f9f9] border-b border-[#e5e5e5]">
                                <span class="w-5 h-5 rounded-full bg-[#888] text-white text-[10px] font-bold flex items-center justify-center">2</span>
                                <span class="text-sm font-semibold text-[#333]">x-api-key header</span>
                        </div>
                        <div class="p-4 bg-[#f4f4f4]">
                                <p class="text-xs text-[#888] mb-3">Alternative for environments that restrict the Authorization header.</p>
                                <CodeBlock code="x-api-key: sk-<your-api-key>" id="auth-xkey" />
                        </div>
                </div>
        </div>

        <!-- Example with SDK -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">SDK examples</h2>
        <CodeBlock
                code={`from openai import OpenAI

client = OpenAI(
    api_key="sk-<your-api-key>",   # ← your key here
    base_url="${$baseUrl || 'https://api.example.com/v1'}"
)

# All requests use the key automatically
response = client.chat.completions.create(
    model="Mistral-Small-4",
    messages=[{"role": "user", "content": "Hello"}]
)`}
                id="auth-sdk-py"
                label="Python — openai SDK"
                lang="python"
        />

        <CodeBlock
                code={`import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-<your-api-key>',   // ← your key here
  baseURL: '${$baseUrl || 'https://api.example.com/v1'}'
});

// All requests use the key automatically
const response = await client.chat.completions.create({
  model: 'Mistral-Small-4',
  messages: [{ role: 'user', content: 'Hello' }]
});`}
                id="auth-sdk-ts"
                label="TypeScript — openai SDK"
                lang="typescript"
        />

        <!-- Warning -->
        <div class="flex items-start gap-3 p-4 rounded-xl border border-amber-200 bg-amber-50 mt-6">
                <svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10.29 3.86L1.82 18a2 2 0 001.71 3h16.94a2 2 0 001.71-3L13.71 3.86a2 2 0 00-3.42 0z"/><line x1="12" y1="9" x2="12" y2="13"/><line x1="12" y1="17" x2="12.01" y2="17"/></svg>
                <div class="text-sm text-amber-800 space-y-1">
                        <p class="font-semibold">Keep your API key secret</p>
                        <p>Never expose your API key in client-side code, public repositories, or logs. Use environment variables or a secrets manager in production.</p>
                </div>
        </div>

        <!-- Key management tip -->
        <div class="mt-6 p-4 rounded-xl border border-[#f0f0f0] bg-[#fafafa]">
                <p class="text-sm font-semibold text-[#333] mb-1">Managing your keys</p>
                <p class="text-sm text-[#666]">
                        Create, rotate, and revoke API keys from the
                        <a href="/" class="text-[#0f0f0f] hover:underline font-medium">Console dashboard</a>.
                        Each key shows usage statistics and can be deactivated instantly.
                </p>
        </div>

</div>
