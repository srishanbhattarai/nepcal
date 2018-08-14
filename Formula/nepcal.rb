class Nepcal < Formula
  desc "Equivalent of Linux's cal, for Nepali dates"
  homepage "https://github.com/nepcal/nepcal"
  url "https://github.com/srishanbhattarai/nepcal/releases/download/v0.4.0/nepcal_0.4.0_darwin_amd64.tar.gz"
  version "0.4.0"
  sha256 "060ef91d1190cae28a219c41a20b6b0bad481baba33c32a3cd91eb682cc2358a"

  def install
    bin.install "nepcal"
  end
end
