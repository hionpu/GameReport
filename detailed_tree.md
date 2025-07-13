# Code Tree Structure

Generated from ctags analysis of the codebase.

**Filtered for subpath:** `dev`

  - 📁 **dev**

    - 📁 **go**

      - 📁 **cmd**

        - 📁 **server**

          - 📄 **main.go**
            - 🔧 main

      - 📁 **internal**

        - 📁 **config**

          - 📄 **config.go**
            - 🏗️ Config
            - 🔧 LoadConfig
            - 🔧 getEnv

        - 📁 **db**

          - 📄 **supabase.go**
            - 🏗️ DB
            - 🏗️ DBConfig
            - 🔧 NewDB

        - 📁 **handlers**

          - 📄 **handlers.go**
            - 🏗️ Handler
            - 📋 webFolderPath
            - 🔧 NewHandler
            - 🔧 loadTemplates

        - 📁 **middleware**

          - 📄 **middleware.go**
            - 🔧 CORS
            - 🔧 Logger
            - 🔧 Recovery
            - 🔧 Timeout

        - 📁 **models**

          - 📄 **models.go**
            - 🏗️ APIResponse
            - 🏗️ DatabaseHealth
            - 🏗️ HealthStatus
            - 🏗️ PageData
            - 🏗️ SupabaseHealth

      - 📁 **web**

        - 📁 **static**

          - 📁 **css**

            - 📄 **custom.css**
              - 🆔 #mobile-menu
              - 🎨 .achievement-glow
              - 🎨 .animate-spin
              - 🎨 .bg-purple-accent
              - 🎨 .bg-purple-bg
              - 🎨 .bg-purple-primary
              - 🎨 .bg-purple-secondary
              - 🎨 .border-purple-100
              - 🎨 .border-purple-200
              - 🎨 .border-purple-primary
              - 🎨 .btn-loading
              - 🎨 .btn-loading::after
              - 🎨 .card-hover
              - 🎨 .card-hover:hover
              - 🎨 .custom-scrollbar
              - 🎨 .custom-scrollbar::-webkit-scrollbar
              - 🎨 .custom-scrollbar::-webkit-scrollbar-thumb
              - 🎨 .custom-scrollbar::-webkit-scrollbar-thumb:hover
              - 🎨 .custom-scrollbar::-webkit-scrollbar-track
              - 🎨 .error-message
              - 🎨 .focus
              - 🎨 .focus
              - 🎨 .from-purple-50
              - 🎨 .from-purple-600
              - 🎨 .hover
              - 🎨 .hover
              - 🎨 .hover
              - 🎨 .htmx-indicator
              - 🎨 .htmx-request .htmx-indicator
              - 🎨 .htmx-request.htmx-indicator
              - 🎨 .htmx-settling
              - 🎨 .mobile-full
              - 🎨 .mobile-stack
              - 🎨 .mobile-text-center
              - 🎨 .progress-bar
              - 🎨 .success-message
              - 🎨 .text-purple-primary
              - 🎨 .text-purple-secondary
              - 🎨 .to-purple-100
              - 🎨 .to-purple-700
              - 🎨 .transition-colors
              - 🏗️ :bg-purple-600:hover
              - 🏗️ :ring-2:focus
              - 🏗️ :ring-purple-primary:focus
              - 🏗️ :text-purple-600:hover
              - 🏗️ :text-purple-primary:hover

        - 📁 **templates**

          - 📁 **components**

            - 📄 **analysis-result.html**
              - 🔌 Analysis Results for {{.SummonerName}}

            - 📄 **sample-report.html**
              - 🔌 Sample Analysis Report

          - 📄 **home.html**
            - 🔌 How It Works
            - 🔌 Powerful Features for Better Gaming
            - 🔌 Ready to Level Up Your Gaming?
            - 🔌 Simple Pricing

    - 📁 **python**

      - 📁 **analyzers**

        - 📄 **base.py**
          - 📦 BaseAnalyzer
            - ⚙️ __init__
            - ⚙️ fetch_data
            - ⚙️ generate_insights
            - ⚙️ get_user_report
            - ⚙️ process_data

        - 📄 **lol_analyzer.py**
          - 📦 LolAnalyzer
            - ⚙️ __init__
            - ⚙️ fetch_data
            - ⚙️ generate_insights
            - ⚙️ get_user_report
            - ⚙️ process_data
            - ⚙️ save_data

      - 📁 **database**

        - 📄 **db_handler.py**
          - 📋 db_handler
          - 📦 DBHandler
            - ⚙️ __init__
            - ⚙️ bulk_insert_matches
            - ⚙️ get_client

      - 📄 **main.py**
        - 📋 app
        - 📋 port
        - 🔧 health_check

