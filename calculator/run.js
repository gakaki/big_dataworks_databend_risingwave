// Databend cost calculator

// Constants
const STORAGE_PRICE_PER_GB_PER_MONTH = 0.019; // USD
const API_CALL_PRICE_PER_MILLION = 0.03; // USD
const CLOUD_SERVICE_FEE_PERCENTAGE = 0.1; // 10%

// Cluster prices per hour (based on Databend pricing)
const clusterPrices = {
    'xs': 0.13,
    's': 0.26,
    'm': 0.52,
    'l': 1.04,
    'xl': 2.08
};

function calculateCosts(clusterSize, storageGB, queriesPerHour) {
    // Storage cost calculation
    const monthlyStorageCost = storageGB * STORAGE_PRICE_PER_GB_PER_MONTH;
    
    // Compute cost calculation
    const activeMinutesPerQuery = 5; // 5 minutes active time per query
    const totalActiveMinutesPerHour = queriesPerHour * activeMinutesPerQuery;
    const activeHoursPerDay = (totalActiveMinutesPerHour / 60) * 24;
    const dailyComputeCost = activeHoursPerDay * clusterPrices[clusterSize];
    
    // API calls calculation
    const apiCallsPerDay = queriesPerHour * 24;
    const apiCallsPerMonth = apiCallsPerDay * 30;
    const apiCallsPerYear = apiCallsPerMonth * 12;
    const monthlyApiCost = (apiCallsPerMonth / 1000000) * API_CALL_PRICE_PER_MILLION;
    
    // Calculate base costs
    const dailyComputeCosts = dailyComputeCost;
    const monthlyComputeCosts = dailyComputeCost * 30;
    const yearlyComputeCosts = monthlyComputeCosts * 12;
    
    const dailyStorageCosts = monthlyStorageCost / 30;
    const yearlyStorageCosts = monthlyStorageCost * 12;
    
    const dailyApiCosts = monthlyApiCost / 30;
    const yearlyApiCosts = monthlyApiCost * 12;
    
    // Calculate cloud service fee
    const dailyCloudFee = (dailyComputeCosts + dailyStorageCosts + dailyApiCosts) * CLOUD_SERVICE_FEE_PERCENTAGE;
    const monthlyCloudFee = (monthlyComputeCosts + monthlyStorageCost + monthlyApiCost) * CLOUD_SERVICE_FEE_PERCENTAGE;
    const yearlyCloudFee = (yearlyComputeCosts + yearlyStorageCosts + yearlyApiCosts) * CLOUD_SERVICE_FEE_PERCENTAGE;
    
    // Calculate totals
    const dailyTotal = dailyComputeCosts + dailyStorageCosts + dailyApiCosts + dailyCloudFee;
    const monthlyTotal = monthlyComputeCosts + monthlyStorageCost + monthlyApiCost + monthlyCloudFee;
    const yearlyTotal = yearlyComputeCosts + yearlyStorageCosts + yearlyApiCosts + yearlyCloudFee;
    
    return {
        period: {
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
        }
    };
}

// Calculate for all cluster sizes
const clusterSizes = ['xs', 's', 'm', 'l', 'xl'];
const storage = 500; // GB
const queriesPerHour = 5;

console.log('Databend Cost Analysis (USD):\n');
clusterSizes.forEach(size => {
    const costs = calculateCosts(size, storage, queriesPerHour);
    console.log(`Cluster Size: ${size.toUpperCase()}`);
    ['daily', 'monthly', 'yearly'].forEach(period => {
        console.log(`\n${period.charAt(0).toUpperCase() + period.slice(1)} Costs:`);
        console.log(`  Compute: $${costs.period[period].compute}`);
        console.log(`  Storage: $${costs.period[period].storage}`);
        console.log(`  API Calls: $${costs.period[period].api}`);
        console.log(`  Cloud Service Fee: $${costs.period[period].cloudFee}`);
        console.log(`  Total: $${costs.period[period].total}`);
    });
    console.log('\n-------------------\n');
});