defmodule GameReport.Models do
  @moduledoc """
  Data structures for GameReport application.
  """

  defmodule HealthStatus do
    @moduledoc """
    Represents the health status of the application.
    """
    defstruct [
      :status,
      :timestamp,
      :version,
      :database,
      :supabase,
      :details
    ]

    @type t :: %__MODULE__{
      status: String.t(),
      timestamp: DateTime.t(),
      version: String.t(),
      database: DatabaseHealth.t(),
      supabase: SupabaseHealth.t(),
      details: map() | nil
    }
  end

  defmodule DatabaseHealth do
    @moduledoc """
    Represents database health information.
    """
    defstruct [
      :status,
      :connections,
      :latency
    ]

    @type t :: %__MODULE__{
      status: String.t(),
      connections: map(),
      latency: String.t()
    }
  end

  defmodule SupabaseHealth do
    @moduledoc """
    Represents Supabase-specific health information.
    """
    defstruct [
      :status,
      :url,
      :connected
    ]

    @type t :: %__MODULE__{
      status: String.t(),
      url: String.t(),
      connected: boolean()
    }
  end

  defmodule PageData do
    @moduledoc """
    Represents data passed to HTML templates.
    """
    defstruct [
      :title,
      :current_date,
      :version,
      :data
    ]

    @type t :: %__MODULE__{
      title: String.t(),
      current_date: String.t(),
      version: String.t(),
      data: map() | nil
    }
  end

  defmodule APIResponse do
    @moduledoc """
    Represents a standard API response.
    """
    defstruct [
      :success,
      :message,
      :data,
      :error
    ]

    @type t :: %__MODULE__{
      success: boolean(),
      message: String.t(),
      data: any(),
      error: String.t() | nil
    }
  end
end