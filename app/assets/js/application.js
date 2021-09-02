require("bootstrap/dist/js/bootstrap.bundle.js");
require("imask/dist/imask");

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
});