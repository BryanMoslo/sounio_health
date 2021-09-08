require("bootstrap/dist/js/bootstrap.bundle.js");
require("imask/dist/imask");
import jsPDFInvoiceTemplate from "jspdf-invoice-template";

$(() => {
    $(document).on('submit','form',function(){
        $(".money-input").each(function (){
            let unmasked = $(this).val().replace(/\D+/g, '')
            $(this).val(unmasked)
        })
    });

    $(".percentage-input").each(function (){
        var numberMask = IMask(this, {
            mask: Number,
            scale: 2,
            signed: false,
            thousandsSeparator: '',
            normalizeZeros: false,
            min: 0,
            max: 99.99,
            padFractionalZeros: true,
            radix: '.'
        }).on('accept', () => {
            this.innerHTML = numberMask.masked.number;
        });
    })

    $(".money-input").each(function (){
        var numberMask = IMask(this, {
            mask: '$num',
            blocks: {
                num: {
                mask: Number,
                thousandsSeparator: ','
                }
            }
        }).on('accept', () => {
            this.innerHTML = numberMask.masked.number;
        });
    })

    $(document).on('change', '.contract-type-select', function() {
        let contractType = $(this).val();
        $(".purchase-option-percentage").addClass("d-none");
        
        if(contractType == "Leasing") {
            $(".purchase-option-percentage").removeClass("d-none");
        }
    })

    $(document).on('change', '#interest-rate-PolicyRatePresent', function() {
        let checked = $(this).is(":checked");
        $(".policy-rate").addClass("d-none");
        
        if(checked) {
            $(".policy-rate").removeClass("d-none");
        }
    })

    $(document).on('click', '.export-pdf', function() {
        let clientName = $(".client-name").html()
        let clientEmail = $(".client-email").html()
        let clientPhoneNumber = $(".client-phone-number").html()
        let clientAddress = $(".client-address").html()

        let equipmentName = $(".equipmen-name").html()
        let equipmentDescription = $(".equipmen-description").html()
        let equipmentPrice = $(".equipment-price").html()
        let contractType = $(".contract-type").html()
        let term = $(".term").html()
        let fee = $(".fee").html()

        let quotationDate = $(".data").data("date")
        let currentDate = $(".data").data("today")
        let policyIncluded = $(".data").data("policy-included") == true
        let policyNotIncludedMessage = "(+ Seguro)"
        if (policyIncluded) {
            policyNotIncludedMessage = ""
        }

        let headers = ["Equipo", "Valor", "Tipo de Contrato", "Plazo", "Cuota Mensual"]
        let values = [
            equipmentName,
            equipmentPrice,
            contractType,
            term,
            `${fee}\n${policyNotIncludedMessage}`
        ]

        if(contractType == "Leasing"){
            let purchaseOptionValue = $(".data").data('purchase-option-value')
            let purchaseOptionPercentage = $(".data").data('purchase-option-percentage')
            headers.push("Opción de Compra")
            values.push(`${purchaseOptionValue}\n(${purchaseOptionPercentage}%)`)
        }

        let quotationForMessage = ""
        if(clientName)  {
            quotationForMessage = "Cotización para:"
        }

        var props = {
            outputType: "",
            returnJsPDFDocObject: true,
            fileName: "Invoice 2021",
            orientationLandscape: false,
            logo: {
                src: "/assets/images/logo.jpg",
                width: 35, //aspect ratio = width/height
                height: 25,
                margin: {
                    top: 0, //negative or positive num, from the current position
                    left: 0 //negative or positive num, from the current position
                }
            },
            business: {
                name: "Sounio Health",
                address: "contacto@souniohealth.com",
                phone: "(+57) 321 827 12 27",
                email: "(+57) 301 796 45 85",
                email_1: "Cra 55 # 187 - 821",
                website: "Cali, Colombia",
            },
            contact: {
                label: quotationForMessage,
                name: clientName,
                address: clientAddress,
                phone: clientPhoneNumber,
                email: clientEmail,
            },
            invoice: {
                label: "Cotización #: ",
                num: Math.floor(Math.random() * (99999 - 999)) + 999,
                invDate: `Fecha de hoy: ${quotationDate}`,
                invGenDate: `Fecha de Cotización: ${currentDate}`,
                headerBorder: false,
                tableBodyBorder: false,
                header: headers,
                table: Array.from(Array(1), (item, index)=>(values)),
                invDescLabel: "Caracteristicas del equipo",
                invDesc: equipmentDescription,
            },
            footer: {
                text: "Gracias por escogernos. Quedamos atentos a sus inquietudes.",
            },
            pageEnable: true,
            pageLabel: "Page ",
        };

        const pdfObject = jsPDFInvoiceTemplate(props); //returns number of pages created

        var pdfCreated = jsPDFInvoiceTemplate.default({});

        pdfCreated.jsPDFDocObject.save();
    })
});