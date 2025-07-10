# Phase 3: 머신러닝(ML) 도입 파이프라인 개요

## 1. 개요 (Overview) 🤖

**목표:** Phase 2까지 축적된 정제된 데이터와 추출된 패턴을 바탕으로, **머신러닝 모델을 학습시켜 인간이 직관적으로 파악하기 어려운 복잡하고 예측적인 인사이트**를 도출한다. '과거'를 분석하는 것을 넘어 '미래'를 예측하고, 초개인화된 피드백을 제공하는 것을 목표로 한다.

## 2. 핵심 기능 명세 (Core Feature Specifications)

- **승패 예측 모델:** 밴픽 단계 또는 게임 15분 시점의 데이터를 기반으로 게임의 승리 확률을 예측한다.
    
- **플레이 스타일 분류 (군집화):** 사용자의 플레이 기록을 '공격적인 초반 개입형', '안정적인 성장형', '운영 중심형' 등과 같이 몇 가지 특징적인 스타일로 자동 분류한다.
    
- **핵심 승리 요인 분석:** 특정 매치업에서 승리에 가장 큰 영향을 미치는 요인(예: '첫 용 획득', '15분 골드 리드')이 무엇인지 모델을 통해 분석한다.
    
- **이상 행동 탐지:** 사용자의 특정 플레이가 평소 자신의 성공 패턴이나 최상위 유저의 패턴과 얼마나 다른지 감지하고 알려준다.
    

## 3. 주요 구현 과업 (Key Implementation Tasks)

1. **ML 데이터셋 구축:**
    
    - Build the training dataset using the clean, aggregated statistics produced by the **high-performance C++ engine from Phase 2**.
    - This ensures the ML models are trained on a robust, large-scale, and efficiently processed dataset. The data will be further refined (feature scaling, one-hot encoding) using Python scripts.
        
2. **모델 선택, 학습 및 평가:**
    
    - **분류 모델:** 승패 예측을 위해 로지스틱 회귀(기본), XGBoost/LightGBM(고성능) 등의 모델을 학습하고 교차 검증을 통해 평가한다.
        
    - **군집화 모델:** 플레이 스타일 분류를 위해 K-Means와 같은 비지도 학습 모델을 적용한다.
        
3. **모델 관리 및 서빙:**
    
    - 학습된 모델을 파일로 저장하고 버전을 관리한다.
        
    - FastAPI를 통해 학습된 모델을 로드하고, 새로운 데이터에 대한 예측 결과를 반환하는 **추론(Inference) API 엔드포인트**를 구축한다.
        
4. **재학습 파이프라인 구축 (MLOps 기초):**
    
    - 게임 메타 변화에 대응하기 위해, 주기적으로 새로운 데이터를 포함하여 모델을 자동으로 재학습하고 성능을 평가하는 파이프라인을 설계한다.
        

## 4. ML 기반 분석 예시

- **(승패 예측):** "조합을 고려했을 때, 이번 게임의 시작 승리 확률은 42%로 다소 불리한 싸움이었습니다."
    
- **(플레이 스타일 분류):** "이번 판은 AI가 분석하기에 '안정적인 성장형' 플레이 스타일에 속합니다. 하지만 이 매치업의 상위 유저들은 보통 '공격적인 초반 개입형'으로 플레이할 때 승률이 더 높았습니다."
    
- **(핵심 요인 분석):** "AI 모델 분석 결과, 당신의 챔피언은 '첫 전령'을 획득했을 때 승률이 25% 상승하는 가장 중요한 승리 요인으로 나타났습니다."


## 4. C++ Engine Integration for ML Pipeline

### C++ Engine Role in ML Feature Engineering

**High-Performance Feature Extraction:**
The C++ engine from Phase 2 becomes crucial for ML model training by providing optimized feature engineering:

```cpp
namespace ml_features {

class FeatureExtractor {
public:
    // Extract features for win prediction models
    std::vector<float> ExtractMatchFeatures(const MatchData& match);
    std::vector<float> ExtractTimelineFeatures(const std::vector<TimelineEvent>& events);
    
    // Extract features for playstyle clustering
    PlayerStyleFeatures ExtractPlayerStyleFeatures(const std::string& puuid, int days = 30);
    
    // Real-time feature computation for inference
    std::vector<float> ComputeRealTimeFeatures(const LiveGameData& game_state);
    
private:
    void ComputeStatisticalFeatures(const MatchData& match, std::vector<float>& features);
    void ComputeTimeseriesFeatures(const std::vector<TimelineEvent>& events, std::vector<float>& features);
    void ComputePerformanceRatios(const MatchData& match, std::vector<float>& features);
};

}
```

### ML Dataset Preparation Performance

**C++ Engine Benefits for ML:**
- **Feature Extraction Speed:** 15x faster than Python pandas operations
- **Memory Efficiency:** 4x less memory usage for large dataset transformations
- **Parallel Processing:** Multi-threaded feature computation for millions of matches
- **Real-time Inference:** Sub-millisecond feature extraction for live predictions

**Dataset Scale Capabilities:**
- **Training Data:** Process 10M+ matches for robust model training
- **Feature Matrix:** Generate 500+ features per match efficiently
- **Cross-validation:** Rapid k-fold dataset splitting and preprocessing
- **Online Learning:** Real-time feature updates for adaptive models

### C++ Engine Enhanced ML Pipeline

```python
# Enhanced ML pipeline with C++ acceleration
class MLPipeline:
    def __init__(self):
        self.cpp_feature_extractor = CppFeatureExtractor()
        self.models = {}
        
    def prepare_training_data(self, match_ids: List[str]):
        # C++ engine extracts features at high speed
        features = self.cpp_feature_extractor.extract_batch_features(match_ids)
        
        # Python handles final transformations and model training
        X_train, y_train = self.preprocess_features(features)
        return X_train, y_train
        
    def real_time_prediction(self, live_game_data):
        # C++ engine computes features in real-time
        features = self.cpp_feature_extractor.compute_realtime_features(live_game_data)
        
        # Python models make predictions
        win_probability = self.models['win_predictor'].predict_proba([features])[0][1]
        return win_probability
```

### Advanced ML Features Enabled by C++ Engine

**1. Real-time Win Prediction:**
- Sub-second feature extraction from live game state
- Continuous model updates as game progresses
- Memory-efficient sliding window computations

**2. Advanced Playstyle Clustering:**
- High-dimensional feature spaces (1000+ features)
- Efficient distance computations for clustering algorithms
- Real-time player classification based on recent matches

**3. Anomaly Detection:**
- Statistical outlier detection in high-dimensional space
- Real-time performance deviation analysis
- Pattern-based unusual behavior identification

**4. Predictive Analytics:**
- Item build outcome prediction with complex feature interactions
- Team composition synergy analysis with combinatorial features
- Meta trend prediction using time-series feature engineering

### Performance Optimization for ML Workloads

**Memory Management:**
- Custom allocators for ML feature matrices
- Memory-mapped file access for large training datasets
- Efficient sparse matrix representations for categorical features

**Compute Optimization:**
- SIMD vectorization for feature calculations
- GPU acceleration for large-scale batch processing (via CUDA interop)
- Parallel pipeline execution for independent feature groups

**Data Pipeline Integration:**
- Zero-copy data transfer between C++ engine and Python ML frameworks
- Streaming data processing for continuous model updates
- Incremental feature computation for online learning scenarios