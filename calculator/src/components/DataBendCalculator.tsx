import React, { useState } from 'react';
import './DatabendCalculator.css';

interface CostDetail {
  compute: string;
  storage: string;
  api: string;
  cloudFee: string;
  total: string;
}

interface CalculationResults {
  daily: CostDetail;
  monthly: CostDetail;
  yearly: CostDetail;
}

// 存储价格：每 TB 每月 160 元（1 TB = 1000 GB）
const STORAGE_PRICE_PER_TB_PER_MONTH = 160;

// 云服务费率
const CLOUD_SERVICE_FEE_PERCENTAGE = 0.1;

// API 调用费用（根据版本选择，每 10,000 次调用费用分别为 ￥3.00 和 ￥4.50，换算后每百万次分别为 300 和 450 元）

// 下面两个对象分别表示基础版和商业版的计算集群价格（单位：人民币/小时）
// 同时根据价格表，每秒费用 = 每小时费用 / 3600
const clusterPricesStandard: Record<string, number> = {
  xs: 3.00,
  s: 6.00,
  m: 12.00,
  l: 24.00,
  xl: 48.00
};

const clusterPricesCommercial: Record<string, number> = {
  xs: 4.50,
  s: 9.00,
  m: 18.00,
  l: 36.00,
  xl: 72.00
};

const DatabendCalculator: React.FC = () => {
  const [storage, setStorage] = useState<number>(500);
  const [queriesPerHour, setQueriesPerHour] = useState<number>(5);
  const [selectedCluster, setSelectedCluster] = useState<string>('xs');
  // 版本选择："standard"代表基础版，"commercial"代表商业版
  const [version, setVersion] = useState<string>('standard');
  // 计费方式选择，"hourly"为按小时计费，"perSecond"为按秒计费
  const [billingMethod, setBillingMethod] = useState<string>('hourly');
  // 新增：每次查询时长（单位：秒），默认 12 秒
  const [queryDuration, setQueryDuration] = useState<number>(12);
  const [results, setResults] = useState<CalculationResults | null>(null);

  const calculateCosts = () => {
    // 根据版本选择 API 每百万次调用费用
    // 标准版：每 10,000 次请求 ￥3.00 => 每百万次请求 ￥3.00 * 100 = 300
    // 商业版：每 10,000 次请求 ￥4.50 => 每百万次请求 ￥4.50 * 100 = 450
    const apiCallPricePerMillion = version === 'standard' ? 300 : 450;

    // 计算存储费用
    // 每月存储费用 = (存储容量 (GB) / 1000) * 每 TB 每月价格
    // 计算存储费用：若存储容量小于1TB，则按1TB收费，否则按实际TB数收费
    const monthlyStorageCost = Math.ceil(storage / 1000) * STORAGE_PRICE_PER_TB_PER_MONTH;
    // 计算计算费用（使用集群的价格）
    // 根据版本选择对应的集群每小时费用
    const hourlyComputeCost =
      version === 'standard'
        ? clusterPricesStandard[selectedCluster]
        : clusterPricesCommercial[selectedCluster];

    let dailyComputeCost = 0;
    if (billingMethod === 'hourly') {
      // 按小时计费模式：使用用户输入的每次查询时长（秒），转换为小时计费
      // 每次查询耗时小时数 = queryDuration / 3600
      // 每日计算费用 = 每小时查询次数 * 24 * (queryDuration/3600) * 每小时费用
      dailyComputeCost = queriesPerHour * 24 * (queryDuration / 3600) * hourlyComputeCost;
    } else {
      // 按秒计费模式：每次查询只收取1秒费用（即使查询时长大于1秒也只收1秒），
      // 每秒费用 = 每小时费用 / 3600
      const perSecondCost = hourlyComputeCost / 3600;
      // 每日计算费用 = 每小时查询次数 * 24 * 1 * perSecondCost
      dailyComputeCost = queriesPerHour * 24 * perSecondCost;
    }
    
    // 计算 API 调用费用
    // 每日 API 调用次数 = 每小时查询次数 * 24
    const apiCallsPerDay = queriesPerHour * 24;
    // 每月 API 调用次数 = 每日 API 调用次数 * 30
    const apiCallsPerMonth = apiCallsPerDay * 30;
    // 每月 API 调用费用 = (每月 API 调用次数 / 1,000,000) * API每百万次调用费用
    const monthlyApiCost = (apiCallsPerMonth / 1000000) * apiCallPricePerMillion;
  
    // 计算周期性费用
    const dailyComputeCosts = dailyComputeCost;
    const monthlyComputeCosts = dailyComputeCost * 30;
    const yearlyComputeCosts = monthlyComputeCosts * 12;
    
    // 存储费用分周期（存储费用以月为单位计算）
    const dailyStorageCosts = monthlyStorageCost / 30;
    const yearlyStorageCosts = monthlyStorageCost * 12;
    
    // API 费用分周期
    const dailyApiCosts = monthlyApiCost / 30;
    const yearlyApiCosts = monthlyApiCost * 12;
    
    // 计算云服务费用
    // 云服务费用 = (计算费用 + 存储费用 + API 调用费用) * 云服务费率
    const dailyCloudFee = (dailyComputeCosts + dailyStorageCosts + dailyApiCosts) * CLOUD_SERVICE_FEE_PERCENTAGE;
    const monthlyCloudFee = (monthlyComputeCosts + monthlyStorageCost + monthlyApiCost) * CLOUD_SERVICE_FEE_PERCENTAGE;
    const yearlyCloudFee = (yearlyComputeCosts + yearlyStorageCosts + yearlyApiCosts) * CLOUD_SERVICE_FEE_PERCENTAGE;
    
    // 汇总各周期总费用
    const dailyTotal = dailyComputeCosts + dailyStorageCosts + dailyApiCosts + dailyCloudFee;
    const monthlyTotal = monthlyComputeCosts + monthlyStorageCost + monthlyApiCost + monthlyCloudFee;
    const yearlyTotal = yearlyComputeCosts + yearlyStorageCosts + yearlyApiCosts + yearlyCloudFee;
  
    // 更新计算结果（保留两位小数）
    setResults({
      daily: {
        compute: dailyComputeCosts.toFixed(2),
        storage: dailyStorageCosts.toFixed(2),
        api: dailyApiCosts.toFixed(2),
        cloudFee: dailyCloudFee.toFixed(2),
        total: dailyTotal.toFixed(2)
      },
      monthly: {
        compute: monthlyComputeCosts.toFixed(2),
        storage: monthlyStorageCost.toFixed(2),
        api: monthlyApiCost.toFixed(2),
        cloudFee: monthlyCloudFee.toFixed(2),
        total: monthlyTotal.toFixed(2)
      },
      yearly: {
        compute: yearlyComputeCosts.toFixed(2),
        storage: yearlyStorageCosts.toFixed(2),
        api: yearlyApiCosts.toFixed(2),
        cloudFee: yearlyCloudFee.toFixed(2),
        total: yearlyTotal.toFixed(2)
      }
    });
  };

  return (
    <div className="max-w-4xl mx-auto p-6 font-sans">
      <h2 className="text-3xl font-bold text-gray-900 mb-8 text-center tracking-tight">
        Databend 成本计算器
      </h2>
      
      <div className="bg-white rounded-xl shadow-lg ring-1 ring-gray-900/5 p-6 mb-8">
        <div className="grid gap-6 mb-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-700">
                存储容量 (GB)
              </label>
              <input
                type="number"
                value={storage}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setStorage(Number(e.target.value))}
                className="w-full px-4 py-2.5 rounded-lg border-0 ring-1 ring-gray-300 
                         text-gray-900 shadow-sm placeholder:text-gray-400
                         focus:ring-2 focus:ring-primary-600 sm:text-sm"
              />
            </div>

            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-700">
                每小时查询次数
              </label>
              <input
                type="number"
                value={queriesPerHour}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setQueriesPerHour(Number(e.target.value))}
                className="w-full px-4 py-2.5 rounded-lg border-0 ring-1 ring-gray-300 
                         text-gray-900 shadow-sm placeholder:text-gray-400
                         focus:ring-2 focus:ring-primary-600 sm:text-sm"
              />
            </div>
          </div>

          {/* 新增：每次查询时间输入（单位：秒） */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-700">
              每次查询时长 (秒)
            </label>
            <input
              type="number"
              value={queryDuration}
              onChange={(e: React.ChangeEvent<HTMLInputElement>) => setQueryDuration(Number(e.target.value))}
              className="w-full px-4 py-2.5 rounded-lg border-0 ring-1 ring-gray-300 
                       text-gray-900 shadow-sm placeholder:text-gray-400
                       focus:ring-2 focus:ring-primary-600 sm:text-sm"
            />
          </div>

          {/* 版本选择 */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-700">
              版本
            </label>
            <select
              value={version}
              onChange={(e: React.ChangeEvent<HTMLSelectElement>) => setVersion(e.target.value)}
              className="w-full px-4 py-2.5 rounded-lg border-0 ring-1 ring-gray-300 
                       text-gray-900 shadow-sm focus:ring-2 focus:ring-primary-600 
                       sm:text-sm"
            >
              <option value="standard">标准版 (每 10,000 次 API 请求 ￥3.00)</option>
              <option value="commercial">商业版 (每 10,000 次 API 请求 ￥4.50)</option>
            </select>
          </div>
          
          {/* 新增：计费方式选择 */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-700">
              计费方式
            </label>
            <select
              value={billingMethod}
              onChange={(e: React.ChangeEvent<HTMLSelectElement>) => setBillingMethod(e.target.value)}
              className="w-full px-4 py-2.5 rounded-lg border-0 ring-1 ring-gray-300 
                       text-gray-900 shadow-sm focus:ring-2 focus:ring-primary-600 
                       sm:text-sm"
            >
              <option value="hourly">按小时计费</option>
              <option value="perSecond">按秒计费 (每次查询仅收1秒费用，后续5分钟内免费)</option>
            </select>
          </div>

          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-700">
              计算集群规格
            </label>
            <select
              value={selectedCluster}
              onChange={(e: React.ChangeEvent<HTMLSelectElement>) => setSelectedCluster(e.target.value)}
              className="w-full px-4 py-2.5 rounded-lg border-0 ring-1 ring-gray-300 
                      text-gray-900 shadow-sm focus:ring-2 focus:ring-primary-600 
                      sm:text-sm"
            >
              <option value="xs">XS (0.13 人民币/小时 | 0.00083333/秒 标准版)</option>
              <option value="s">S (0.26 人民币/小时 | 0.00166667/秒 标准版)</option>
              <option value="m">M (0.52 人民币/小时 | 0.00333333/秒 标准版)</option>
              <option value="l">L (1.04 人民币/小时 | 0.00666667/秒 标准版)</option>
              <option value="xl">XL (2.08 人民币/小时 | 0.01333333/秒 标准版)</option>
            </select>
          </div>

          <button
            onClick={calculateCosts}
            className="w-full bg-primary-600 text-white py-3 rounded-lg 
                     hover:bg-primary-700 focus:outline-none focus:ring-2 
                     focus:ring-offset-2 focus:ring-primary-600 
                     transition duration-200 ease-in-out"
          >
            计算成本
          </button>
        </div>
      </div>

      {results && (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {(['daily', 'monthly', 'yearly'] as const).map((period) => (
            <div key={period} 
                 className="bg-white rounded-xl shadow-lg ring-1 ring-gray-900/5 p-6 
                          hover:shadow-xl transition duration-300">
              <h3 className="text-xl font-semibold text-gray-900 mb-4 text-center">
                {period === 'daily'
                  ? '每日'
                  : period === 'monthly'
                  ? '每月'
                  : '每年'}成本
              </h3>
              
              <div className="space-y-3">
                <div className="flex justify-between items-center text-sm">
                  <span className="text-gray-600">计算费用</span>
                  <span className="font-medium text-gray-900">¥{results[period].compute}</span>
                </div>
                <div className="flex justify-between items-center text-sm">
                  <span className="text-gray-600">存储费用</span>
                  <span className="font-medium text-gray-900">¥{results[period].storage}</span>
                </div>
                <div className="flex justify-between items-center text-sm">
                  <span className="text-gray-600">API调用费用</span>
                  <span className="font-medium text-gray-900">¥{results[period].api}</span>
                </div>
                <div className="flex justify-between items-center text-sm">
                  <span className="text-gray-600">云服务费用</span>
                  <span className="font-medium text-gray-900">¥{results[period].cloudFee}</span>
                </div>
                
                <div className="pt-3 mt-3 border-t border-gray-200">
                  <div className="flex justify-between items-center">
                    <span className="font-semibold text-gray-900">总计</span>
                    <span className="font-bold text-lg text-primary-600">
                      ¥{results[period].total}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default DatabendCalculator;