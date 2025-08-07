defmodule GameReport.Config do
  @moduledoc """
  Configuration module for GameReport application.
  """

  def port, do: System.get_env("PORT", "8080")
  
  def database_url do
    System.get_env("DATABASE_URL", "postgresql://postgres:ckathwn2%40@db.fssbljnxonqzwctasvjk.supabase.co:5432/postgres")
  end
  
  def supabase_url do
    System.get_env("SUPABASE_URL", "https://fssbljnxonqzwctasvjk.supabase.co")
  end
  
  def supabase_anon_key do
    System.get_env("SUPABASE_ANON_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZzc2Jsam54b25xendjdGFzdmprIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTE4NjYxNjksImV4cCI6MjA2NzQ0MjE2OX0._ecGHtXHup28xAW6svhlVZw4LUzCzQj1vVzxoud5_I4")
  end
  
  def supabase_service_role_key do
    System.get_env("SUPABASE_SERVICE_ROLE_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZzc2Jsam54b25xendjdGFzdmprIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc1MTg2NjE2OSwiZXhwIjoyMDY3NDQyMTY5fQ.UHfkcjGAjG5I269PQPLYyKNs9dsLXWYBApbTCZX-ygk")
  end
  
  def environment, do: System.get_env("ELIXIR_ENV", "development")
  
  def validate! do
    if String.trim(database_url()) == "" do
      raise "DATABASE_URL is required"
    end
    
    if String.trim(supabase_url()) == "" do
      raise "SUPABASE_URL is required"
    end
    
    :ok
  end
end