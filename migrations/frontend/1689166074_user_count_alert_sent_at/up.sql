ALTER TABLE product_licenses ADD COLUMN IF NOT EXISTS user_count_alert_sent_at TIMESTAMP WITH TIME ZONE;
