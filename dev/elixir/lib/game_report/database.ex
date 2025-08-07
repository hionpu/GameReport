defmodule GameReport.Database do
  @moduledoc """
  Database connection and utilities for GameReport application.
  """

  use GenServer
  require Logger

  @connection_timeout 5000
  @default_pool_size 10

  def start_link(_) do
    GenServer.start_link(__MODULE__, [], name: __MODULE__)
  end

  def init(_) do
    database_url = Application.get_env(:game_report, GameReport.Database)[:database_url]
    
    # Parse database URL into connection options
    uri = URI.parse(database_url)
    [user, password] = String.split(uri.userinfo || "", ":")
    
    connection_config = [
      hostname: uri.host,
      port: uri.port,
      database: String.trim_leading(uri.path || "", "/"),
      username: URI.decode(user || "postgres"),
      password: URI.decode(password || ""),
      pool_size: @default_pool_size,
      timeout: @connection_timeout,
      backoff_type: :exp,
      backoff_min: 64,
      backoff_max: 10_000
    ]

    case Postgrex.start_link(connection_config) do
      {:ok, pid} ->
        Logger.info("🔗 Database connection established")
        {:ok, %{connection: pid, stats: %{connections: @default_pool_size}}}
      
      {:error, reason} ->
        Logger.error("❌ Failed to connect to database: #{inspect(reason)}")
        {:ok, %{connection: nil, stats: %{connections: 0}}}
    end
  end

  def health_check do
    case GenServer.call(__MODULE__, :health_check) do
      {:ok, _result} -> :ok
      {:error, reason} -> {:error, reason}
    end
  end

  def get_statistics do
    GenServer.call(__MODULE__, :get_stats)
  end

  def handle_call(:health_check, _from, %{connection: nil} = state) do
    {:reply, {:error, "No database connection"}, state}
  end

  def handle_call(:health_check, _from, %{connection: conn} = state) do
    try do
      case Postgrex.query(conn, "SELECT 1", []) do
        {:ok, _} -> {:reply, {:ok, "Database healthy"}, state}
        {:error, reason} -> {:reply, {:error, reason}, state}
      end
    rescue
      e -> {:reply, {:error, e}, state}
    end
  end

  def handle_call(:get_stats, _from, %{stats: stats} = state) do
    enhanced_stats = Map.put(stats, :timestamp, DateTime.utc_now())
    {:reply, enhanced_stats, state}
  end
end