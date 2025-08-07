# This file is responsible for configuring your application
# and its dependencies with the aid of the Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
import Config

config :game_report,
  generators: [timestamp_type: :utc_datetime]

# Database configuration
config :game_report, GameReport.Database,
  database_url: System.get_env("DATABASE_URL", "postgresql://postgres:ckathwn2%40@db.fssbljnxonqzwctasvjk.supabase.co:5432/postgres")

# Supabase configuration
config :game_report, GameReport.Supabase,
  url: System.get_env("SUPABASE_URL", "https://fssbljnxonqzwctasvjk.supabase.co"),
  anon_key: System.get_env("SUPABASE_ANON_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZzc2Jsam54b25xendjdGFzdmprIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTE4NjYxNjksImV4cCI6MjA2NzQ0MjE2OX0._ecGHtXHup28xAW6svhlVZw4LUzCzQj1vVzxoud5_I4"),
  service_role_key: System.get_env("SUPABASE_SERVICE_ROLE_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZzc2Jsam54b25xendjdGFzdmprIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc1MTg2NjE2OSwiZXhwIjoyMDY3NDQyMTY5fQ.UHfkcjGAjG5I269PQPLYyKNs9dsLXWYBApbTCZX-ygk")

# Configures the endpoint
config :game_report, GameReportWeb.Endpoint,
  url: [host: "localhost"],
  adapter: Bandit.PhoenixAdapter,
  render_errors: [
    formats: [html: GameReportWeb.ErrorHTML, json: GameReportWeb.ErrorJSON],
    layout: false
  ],
  pubsub_server: GameReport.PubSub,
  live_view: [signing_salt: "DTF/LqGa"]

# Configures the mailer
#
# By default it uses the "Local" adapter which stores the emails
# locally. You can see the emails in your browser, at "/dev/mailbox".
#
# For production it's recommended to configure a different adapter
# at the `config/runtime.exs`.
config :game_report, GameReport.Mailer, adapter: Swoosh.Adapters.Local

# Configure esbuild (the version is required)
config :esbuild,
  version: "0.17.11",
  game_report: [
    args:
      ~w(js/app.js --bundle --target=es2017 --outdir=../priv/static/assets --external:/fonts/* --external:/images/*),
    cd: Path.expand("../assets", __DIR__),
    env: %{"NODE_PATH" => Path.expand("../deps", __DIR__)}
  ]

# Configure tailwind (the version is required)
config :tailwind,
  version: "3.4.3",
  game_report: [
    args: ~w(
      --config=tailwind.config.js
      --input=css/app.css
      --output=../priv/static/assets/app.css
    ),
    cd: Path.expand("../assets", __DIR__)
  ]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{config_env()}.exs"
