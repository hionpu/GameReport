defmodule GameReportWeb.PageController do
  use GameReportWeb, :controller
  
  alias GameReport.{Config, Database, Models}
  require Logger

  def home(conn, _params) do
    Logger.info("Home page request received from #{get_remote_ip(conn)}")

    page_data = %Models.PageData{
      title: "GameReport",
      current_date: DateTime.utc_now() |> Calendar.strftime("%Y년 %m월 %d일 %A"),
      version: "1.0.0, phase 1",
      data: %{
        environment: Config.environment(),
        supabase_url: Config.supabase_url(),
        features: [
          "Daily Gaming Insights",
          "Performance Metrics",
          "Progress Tracking(Supported later)",
          "AI-Powered Analysis(Supported later)"
        ]
      }
    }

    render(conn, :home, page_data: page_data)
  end

  def health(conn, _params) do
    Logger.info("Health check request received from #{get_remote_ip(conn)}")

    db_start = System.monotonic_time()
    db_result = Database.health_check()
    db_latency = System.monotonic_time() - db_start
    db_latency_ms = System.convert_time_unit(db_latency, :native, :millisecond)

    database_health = %Models.DatabaseHealth{
      status: if(db_result == :ok, do: "OK", else: "BAD"),
      connections: Database.get_statistics(),
      latency: "#{db_latency_ms}ms"
    }

    supabase_health = %Models.SupabaseHealth{
      status: "OK",
      url: Config.supabase_url(),
      connected: true
    }

    health_status = %Models.HealthStatus{
      status: if(db_result == :ok, do: "OK", else: "BAD"),
      timestamp: DateTime.utc_now(),
      version: "1.0.0, phase 1",
      database: database_health,
      supabase: supabase_health,
      details: if(db_result != :ok, do: %{database_error: inspect(db_result)}, else: nil)
    }

    status_code = if(db_result == :ok, do: 200, else: 503)

    if db_result == :ok do
      Logger.info("✅ Health check passed in #{db_latency_ms}ms")
    else
      Logger.error("❌ Health check failed: #{inspect(db_result)}")
    end

    conn
    |> put_status(status_code)
    |> json(health_status)
  end

  def not_found(conn, _params) do
    method = conn.method
    path = conn.request_path
    Logger.info("404 Not Found: #{method} #{path}")

    response = %Models.APIResponse{
      success: false,
      message: "Endpoint not found",
      error: "The requested endpoint #{method} #{path} does not exist"
    }

    conn
    |> put_status(:not_found)
    |> json(response)
  end

  defp get_remote_ip(conn) do
    case Plug.Conn.get_req_header(conn, "x-forwarded-for") do
      [ip | _] -> ip
      [] -> to_string(:inet_parse.ntoa(conn.remote_ip))
    end
  end
end
