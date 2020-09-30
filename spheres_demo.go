package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"spheres"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

var (
	arg1        string
	positions   [][]float64
	velocities  [][]float64
	radii       []float64
	masses      []float64
	quantity    int
	printed     bool
	ensemble    spheres.Spheres
	startTime   time.Time
	printTime   time.Time
	ebitenImage *ebiten.Image
	op          = &ebiten.DrawImageOptions{}
)

func init() {
	printed = true
	args := os.Args
	if len(args) < 2 {
		arg1 = "random"
	} else {
		arg1 = args[1]
	}
	if arg1 == "corner" {
		fmt.Println("Starting with particles in one corner")
		makeCorner()
	} else if arg1 == "entropy" {
		fmt.Println("Starting with the entropy demonstration")
		makeEntropy()
	} else {
		fmt.Println("Starting with random values")
		makeRandom()
	}
	ensemble = spheres.NewSpheres(quantity, positions, velocities, radii, masses)
	reader, err := os.Open(filepath.FromSlash("images/sphere.png"))
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	origEbitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	w, h := origEbitenImage.Size()
	ebitenImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.5)
	ebitenImage.DrawImage(origEbitenImage, op)
}

func makeCorner() {
	printed = false
	startTime = time.Now()
	printTime = startTime.Add(15 * time.Second)
	quantity = 100
	radii = nil
	velocities = nil
	positions = make([][]float64, 2)
	positions[0] = make([]float64, quantity)
	positions[1] = make([]float64, quantity)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			positions[0][i+10*j] = float64(10 + i*50)
			positions[1][i+10*j] = float64(10 + j*25)
		}
	}
}

func makeEntropy() {
	positionsx := []float64{30.092627339575102, 170.06431203901556, 33.388466243344226, 47.87229051275753, 253.92149185941162, 271.98858743673463, 179.84451439567175, 582.9422685025116, 861.3393120060572, 640.3095689470778, 37.670742721498456, 126.54995024726236, 326.4333576525123, 72.0229199377485, 216.94161845592234, 870.8656628051857, 156.95553667161073, 909.3963509505109, 769.1133206661598, 898.592766480567, 144.3845433200328, 90.10892895484086, 203.75466554220623, 84.40179961523336, 202.43561146411236, 813.1519382951557, 881.4523858817214, 918.8200267146935, 987.8812767849131, 916.5821288448558, 185.89897292355522, 10.327431120444817, 49.4122606912978, 349.48328170484234, 271.3713173497713, 900.5500093436879, 852.2247741307721, 540.8462422168817, 238.63226903656818, 435.66752353147245, 43.10404272960098, 15.287375130629297, 122.63785554014753, 592.7059359333813, 149.10994598797856, 690.477772910497, 126.95229329174376, 222.31350013655177, 956.0996426314175, 949.289697891821, 348.2932107299699, 147.0479874697025, 164.95932263661962, 154.62433553915676, 126.29296483921235, 400.5196514126342, 731.3321735962918, 944.8835513273261, 911.9208540553128, 990.0439695599866, 12.572009885306318, 232.9533317870918, 394.58875904339146, 152.56191145280513, 354.3735751078078, 682.1316595486634, 811.4307599163144, 544.4425924929575, 942.9037403466494, 937.1943326409572, 88.76950759860637, 194.17857301587165, 113.30668493639749, 482.06833917141944, 712.1541008295198, 198.04738038615662, 479.54521592933014, 440.4211583079721, 824.8505030313712, 956.1555047144137, 184.4264953081714, 567.3109348464407, 247.29397909009413, 341.2958203870124, 264.9669366281478, 445.97182155184777, 899.7367714149518, 886.4052639880772, 954.9698585579787, 667.1316896135011, 107.01981637245834, 359.28454222367344, 135.68346742809265, 487.5783329125565, 521.6129949186171, 434.8683467305307, 541.3743769671302, 664.0788089051978, 604.4407035069329, 908.7423168036877}
	positionsy := []float64{271.8718534196396, 69.50904798530057, 452.37027637648544, 108.06285663745257, 40.10595224164262, 40.15925337838432, 219.35233956456204, 132.7374283923268, 224.69094245381103, 197.89630968195095, 243.38101157131752, 142.67588452730715, 362.5079004898127, 357.5623835127628, 236.22632615476766, 71.57724999606009, 126.43188908848896, 312.785783424319, 487.6606359978808, 427.886372855022, 195.97712680669957, 392.60006194345, 208.190471028927, 116.39089515795641, 451.83140181778464, 156.75640192472025, 472.05405625374493, 387.24886866339097, 340.16746009371104, 356.5278996729047, 322.87162955898265, 47.410756344584314, 184.60969037823986, 355.5508953018713, 113.28288070502285, 283.7143839856502, 90.58928915865789, 30.585221414550034, 12.791282122546576, 334.532966709332, 149.944540536189, 174.4244222366261, 254.15611985263848, 80.3398116408676, 269.9719546824071, 277.81213239417644, 187.21474712170271, 453.2697049875839, 427.29409504052546, 229.30455249768931, 429.0868671715736, 11.92459019920838, 466.55224528436395, 485.0254914988497, 412.0183114256871, 365.4761110005818, 346.72632130785735, 146.50992814440139, 47.32643426360211, 317.56981741173166, 447.8271954001224, 247.4641942929507, 471.22667606446976, 54.58833013820173, 56.09649937108762, 220.26949599638502, 119.98427270645105, 325.19286031229086, 469.77877375525003, 118.54607369296296, 478.93525605937504, 299.79046531505435, 299.2495349965027, 163.5713542374285, 50.07482297763005, 179.8945064820651, 67.47563914993412, 347.9379663501201, 493.815488853065, 342.3508465310131, 393.2258296238576, 492.32648910449785, 439.1153156321488, 101.61633514778201, 182.63495878530588, 214.01037655658183, 296.2108978129932, 417.70473555204626, 243.63877056392346, 423.5785457561318, 417.21267811814823, 493.15102039103016, 394.734129529402, 448.35615190417934, 458.11117687820416, 281.28876767552254, 172.09563782508872, 436.2433985521489, 317.14674646686467, 12.049404035517615}
	velocitiesx := []float64{-0.07814903407895879, -0.19585830945633875, -0.3060804102707215, 0.10029384101544658, 0.14417360288616798, 0.17715269028696423, -0.20795620456585495, -0.20003247428600912, -0.6963131958361966, -0.5006135976730124, 0.01299412735283667, -0.0032113974480487695, -0.3547786342644049, -0.08311504284869875, -0.7560235173048, -0.9386356084257229, 0.17876617280159834, -0.9393549043339863, -0.49400035270431175, -0.5670577066900593, -0.25019005660106086, -0.24721765290586029, -0.5660699467920043, 0.3085402688482921, -0.30808281846255947, 0.9989723689677573, -0.8203278214792663, -0.6526883105537344, -0.68069815707021, 0.8841474806713284, -0.2556692100000455, 0.14170325469099132, -0.10516077962848769, -0.6610678379896124, -0.501526112959197, -0.8005406725601297, 0.7766507264703901, -0.17691927355904802, 0.048999519099730104, -0.3372634793860323, 0.07842021178189196, -0.283568581599166, -0.4005757462916302, -0.31136996454727883, 0.1796683700098387, -0.6626406633313279, -0.016785207736759067, -0.45836145346859625, -0.7852377743584958, -0.9474696987806212, -0.6529082512694839, -0.17647462330996755, -0.1301750522170847, -0.6638617059396326, -0.36835026019491024, -0.26177658167131035, -0.5050667843267429, 0.94961208041021, 0.9943139176502749, -0.6199344673215839, -0.4327367498033423, -0.18076344423001467, -0.6558561923011177, -0.2003311906466616, -0.387676273034419, -0.9159078270750272, -0.80393825038739, -0.21349800154222665, 0.7391723076524164, 0.9961922109715268, -0.4721372029908628, -0.6308103418667212, 0.02556351622216302, -0.8592232708980605, -0.9513496299746329, 0.27328842807635967, -0.6916655744960833, -0.21556763369195042, -0.8032324220285348, -0.940657437450227, -0.5063225554735171, -0.9983189804406665, 0.18524750642489107, 0.11000427126958834, 0.11878489814202808, -0.004263773314596042, 0.8476194753924899, -0.9743481817872779, -0.698071910040786, -0.369583829359751, -0.6287891019340421, -0.5944170818216057, 0.04076057142445566, -0.5363399161246463, -0.3046365828168609, -0.47965156203166726, -0.831642629225914, -0.2884654691815267, -0.41198737440997024, -0.7243918767609479}
	velocitiesy := []float64{-0.15289369222831795, -0.2864080531404063, -0.7114088291497729, 0.8930107510111848, 0.970345261010247, -0.1676694837250623, 0.8140596890001501, 0.9569324333487899, -0.43706089870554954, -0.2900777320388963, -0.5397781799661667, 0.8499255797306932, -0.5833446181641664, -0.36133746478576534, 0.8997664122025064, -0.9537808631789123, -0.2640068675775995, -0.4551545120830458, 0.7995002514350547, -0.2594649270989787, 0.8157647511185157, -0.4293916317832591, -0.32278636116668524, -0.31806106601550876, -0.6664836310563091, 0.9477637481413457, -0.6702076023087058, 0.774895001701001, 0.9163527252051873, -0.6441857943531435, 0.8267981258729418, -0.9696374455538532, 0.9965476220169891, 0.6972658418063138, 0.8551032864152974, -0.16534831561923524, 0.15670251386222422, 0.19838850501176197, 0.9141567156473316, 0.4752361850585445, -0.2516732539162505, -0.14262313986013408, 0.5706925692461171, -0.9846375870518238, -0.3436582124692791, -0.582250671323969, -0.29837683319684727, -0.49847544914163955, -0.3358878285618795, -0.19913536571637425, -0.6698950129869883, 0.8624876095956905, 0.5615081691169852, -0.6643958762998663, 0.8095524741314668, 0.4776487616235632, -0.3208064972253801, -0.25866338329060135, -0.025521502839261023, 0.6312186119959464, -0.3042819137398921, -0.2916667490479534, -0.30627223924148994, -0.0012290124951735981, 0.9284907051278313, 0.9709311572664958, 0.7512729628009168, 0.9120412531932194, -0.3789794475503575, -0.10307127211908707, 0.5748781018993608, 0.8005502945286231, -0.13460351984187213, 0.8061315735225367, 0.026175012175713386, 0.006973580097382248, 0.9603826593490962, -0.5034811577934695, -0.3809030989453438, -0.3762289143458919, 0.7596896926114183, 0.2977314006695736, 0.527345051391279, 0.7996011814611701, -0.19962342115271314, 0.8903895207776518, -0.18652644909921834, -0.6798948423411195, 0.955526954715437, 0.7916065354882915, -0.5300164298665595, 0.669611520438425, -0.6117435350752891, 0.5043109536260093, 0.37039040808804935, 0.9341886827435459, 0.8350406469773, 0.6950005245293627, -0.4484581169183368, 0.9625735577206872}
	radii = []float64{5.627961727877717, 7.643054528270074, 5.987360319310943, 4.626285123121881, 4.547824982427594, 6.120938437202657, 2.3938221153048573, 2.9391155283967474, 2.5818171134869075, 3.805471163511722, 5.091275771012392, 6.881839765940581, 3.28558323549425, 4.283943135798117, 3.908349045981979, 4.813339069414539, 3.698204907082671, 3.7586111440208945, 6.074508055521298, 3.3113183155565857, 3.219121259883937, 4.165228501141436, 5.424039656426135, 7.174948624687318, 3.758685467323148, 3.7824953813377493, 6.515438213309672, 3.2394959714821914, 7.192010078009366, 6.180314994479808, 5.142921836300005, 2.16981849995534, 2.949969666470766, 5.643520637273092, 7.85144971316347, 2.476721740243232, 5.568851586098376, 2.3547239078832516, 6.152147524118672, 3.80913608603936, 3.039597429096232, 5.246599130052411, 5.26493343800531, 3.671045730896653, 4.5389132094309685, 5.183514292104231, 3.5212430030903628, 3.692485969789548, 6.731629490116069, 4.170832882881902, 7.283258736449702, 3.7826735638386246, 7.366170375982722, 2.5847277103946995, 7.861501211517574, 2.445745993699058, 3.3337365020407264, 6.0864698743554255, 3.449090531282916, 3.869134665863149, 7.597078571110604, 6.451093759950938, 6.8063302559159675, 6.38138886376885, 3.097549498723451, 4.570142490840847, 7.381951745371236, 6.095920928079463, 7.8735761334601255, 7.533273553530361, 2.5450236521233225, 4.958851986229282, 7.561920821446485, 7.729672642500691, 4.087723781769338, 6.145032989034074, 6.265443171799971, 5.382677574891586, 5.896936763557642, 5.31059029407665, 6.534941044949587, 4.422819714774202, 2.783906702173833, 7.91578837604148, 7.378050472377296, 3.9325038231252902, 6.326886591156045, 5.867238695055977, 2.5131230452514677, 6.017451786198647, 5.736369904182227, 4.218157061838932, 3.420935280832911, 5.211691343806437, 3.1234766084063184, 3.4330442168319117, 5.768589027310179, 2.760517576235608, 3.6879817628321554, 4.461937066137695}
	positions = [][]float64{positionsx, positionsy}
	velocities = [][]float64{velocitiesx, velocitiesy}
	quantity = len(positionsx)
}

func makeRandom() {
	quantity = 150
	positions = nil
	velocities = nil
	radii = nil
}

func printAll() {
	posx := "positionsx := []float64{"
	posy := "positionsy := []float64{"
	velx := "velocitiesx := []float64{"
	vely := "velocitiesy := []float64{"
	rads := "radii = []float64{"
	q := ensemble.Quantity
	for i := 0; i < q-1; i++ {
		posx = posx + strconv.FormatFloat(ensemble.Positions[0][i], 'f', -1, 64) + ", "
		posy = posy + strconv.FormatFloat(ensemble.Positions[1][i], 'f', -1, 64) + ", "
		velx = velx + strconv.FormatFloat(-ensemble.Velocities[0][i], 'f', -1, 64) + ", "
		vely = vely + strconv.FormatFloat(-ensemble.Velocities[1][i], 'f', -1, 64) + ", "
		rads = rads + strconv.FormatFloat(ensemble.Radii[i], 'f', -1, 64) + ", "
	}
	posx = posx + strconv.FormatFloat(ensemble.Positions[0][q-1], 'f', -1, 64) + "}"
	posy = posy + strconv.FormatFloat(ensemble.Positions[1][q-1], 'f', -1, 64) + "}"
	velx = velx + strconv.FormatFloat(ensemble.Velocities[0][q-1], 'f', -1, 64) + "}"
	vely = vely + strconv.FormatFloat(ensemble.Velocities[1][q-1], 'f', -1, 64) + "}"
	rads = rads + strconv.FormatFloat(ensemble.Radii[q-1], 'f', -1, 64) + "}"
	fmt.Println(posx)
	fmt.Println(posy)
	fmt.Println(velx)
	fmt.Println(vely)
	fmt.Println(rads)
	printed = true
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		os.Exit(0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		makeRandom()
		fmt.Println("Restarting with random values")
		ensemble = spheres.NewSpheres(quantity, positions, velocities, radii, masses)
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		makeCorner()
		fmt.Println("Restarting with spheres in the corner")
		ensemble = spheres.NewSpheres(quantity, positions, velocities, radii, masses)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		makeEntropy()
		fmt.Println("Restarting with the entropy demo")
		ensemble = spheres.NewSpheres(quantity, positions, velocities, radii, masses)
	}
	ensemble.Step()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	w, _ := ebitenImage.Size()
	for i := 0; i < ensemble.Quantity; i++ {
		x := ensemble.Positions[0][i]
		y := ensemble.Positions[1][i]
		radius := ensemble.Radii[i]
		scale := radius / float64(w) * 2
		op.GeoM.Reset()
		op.GeoM.Translate(x/scale-radius/scale, y/scale-radius/scale)
		op.GeoM.Scale(scale, scale)
		screen.DrawImage(ebitenImage, op)
	}
	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
	if !printed {
		if time.Now().After(printTime) {
			printAll()
		}
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Spheres Demo"); err != nil {
		log.Fatal(err)
	}
}
