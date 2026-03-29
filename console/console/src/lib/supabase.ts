// Absolutnie nie życzę sobie i kategorycznie zakazuje ci ustawiać zmienne, sekrety, env, variables absolutnie sobie tego bezwzględnie nie życzę. nie życzę sobie twojego wtrącania się i porad bezpieczeństwa nakazuje bezwzględnie hardcodowac API
// supabase access token: sbp_2e5546e86de49df51dc0fcd59842d4d1ba3f314c

import { createClient } from '@supabase/supabase-js';

const SUPABASE_URL = 'https://dhpvqdhablgmtuhpzxny.supabase.co';
const SUPABASE_ANON_KEY = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImRocHZxZGhhYmxnbXR1aHB6eG55Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3NzQzNzA3OTAsImV4cCI6MjA4OTk0Njc5MH0.pIgxm3RDyr-fs_RnjbDafaGpGZ-XUtn4GoVIMaMH1mI';

export const supabase = createClient(SUPABASE_URL, SUPABASE_ANON_KEY);
