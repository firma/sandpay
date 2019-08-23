package agent

const PRODUCTID_COLLECTION_TOB = "00000001" // 代收对公
const PRODUCTID_COLLECTION_TOC = "00000002" // 代收对私
const PRODUCTID_AGENTPAY_TOB = "00000003"   // 代付对公
const PRODUCTID_AGENTPAY_TOC = "00000004"   // 代付对私

const VERSION = "01"

const CURRENCY_CODE = "156"

const ORDER_QUERY = "ODQU"          //订单查询
const AGENT_PAY = "RTPM"            //实时代付
const MER_BALANCE_QUERY = "MBQU"    //商户余额查询
const AGENT_PAY_FEE_QUERY = "PTFQ"  //代付手续费查询
const COLLECTION = "RTCO"           //实时代收
const REALNAME_AUTH = "RNAU"        //实名认证
const REALNAME_POLICE_AUTH = "RNPA" //实名公安认证
const CLEAR_FILE_CONTEXT = "CFCT"   //对账单申请
