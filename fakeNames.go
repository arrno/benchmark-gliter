package main

import "strings"

var nameStr string = `Nexora Systems
ByteForge Solutions
CloudSync Innovations
CyberNest Technologies
QuantumByte Labs
CodeHorizon Inc.
TitanSoft Solutions
ApexNova AI
InfinityWare
NeuralGrid
Zenith Capital Group
Pinnacle Wealth Advisors
LedgerLine Consulting
FusionBank Financials
Silvercrest Capital
BluePeak Investments
ClarityEdge Consulting
SummitPoint Strategies
VectorTrust Solutions
AlphaPrime Financial
TrendVana
ShopEase Global
UrbanHive Outfitters
NovaCart Marketplace
LuxeLane Retail
OmniTrend
SnapBuy Direct
VeloxMart
BazaarBlitz
ModernEdge Merchants
MediSphere Solutions
BioNexa Labs
Vitalis Medical
LifeBloom Technologies
Genevia Health
CuraNova Systems
HealSync Innovations
MedElite Diagnostics
OmniCare Wellness
PulseLine Biotech
VeloDrive Motors
Quantum AutoWorks
TurboTrack Logistics
HyperLane Transit
ApexRides Mobility
NitroFleet Solutions
NovaMotion Auto
Zenith Transport Group
DriveSphere Solutions
AutoGlide Technologies
EcoNova Solutions
GreenWave Energy
SunFusion Power
TerraVolt Renewables
PureElement Resources
HydroGenix Systems
BrightLeaf Solar
EnviroSphere Technologies
NeoTerra Energy
BlueHorizon Renewables
StarForge Studios
Vortex Media Group
EchoVerse Productions
SonicPulse Records
HyperLens Entertainment
InfinityWave Films
Zenith Broadcasting
NovaFrame Pictures
PrimeStage Media
CloudNine Animation
UrbanHarvest Foods
ZenithBrew Coffee
PureNest Organics
InfusionBites
TerraFlame Grills
FreshSphere Foods
GoldenCrust Bakery
OmniTaste Catering
SavorEdge Beverages
CulinaryCrest
LuxeEdge Clothing
TrendNova Wear
EliteFabric Designs
ApexStyle Apparel
ModeSphere Fashions
UrbanStitch Collective
RadianceWear
ZenithAttire
NovaThread Designs
IconicWeave
HyperHive Ventures
SummitEdge Labs
OmniXcelerate
StratosFlow Solutions
NovaGrid Innovations
VortexSync Enterprises
ZenithSphere Inc.
HorizonNexus
QuantumLeap Solutions
AlphaCore Innovations`

var nameList []string

func names() []string {
	if nameList != nil {
		return nameList
	}
	nameList = strings.Split(nameStr, "\n")
	return nameList
}
